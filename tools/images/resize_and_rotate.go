package main

import (
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

var wg sync.WaitGroup

const maxFileSize = 5 * 1024 * 1024
const destPath = "out/"

// Computer display standard https://en.wikipedia.org/wiki/Computer_display_standard
var imageBoundSteps = []int{7680, 6400, 5120, 4096, 3840, 3440, 3200, 2880, 2560, 2048, 1920}
var supportedFormats = []string{"jpg", "jpeg", "png"}

type source struct {
	name string
	path string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Source path is empty. Please spcesify files or folders to process")
	}

	var sourceList = getSourceList(os.Args[1:])
	sourceList = copyAll(destPath, sourceList)

	wg.Add(len(sourceList))
	for _, source := range sourceList {
		go process(source.path)
	}
	wg.Wait()

	log.Println("All Done")

}

func copyAll(dstDirPath string, sourceList []source) []source {
	os.RemoveAll(dstDirPath)
	err := os.Mkdir(dstDirPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	for i, source := range sourceList {
		var dstPath = dstDirPath + source.name
		err := copyFile(source.path, dstPath)
		if err != nil {
			log.Fatal(err)
		}
		sourceList[i].path = dstPath
	}
	return sourceList
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func getSourceList(args []string) []source {
	var sourceList = []source{}
	for _, filePath := range args {
		info, err := os.Stat(filePath)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.IsDir() {
			files, err := ioutil.ReadDir(filePath)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				var fileInDirPath = filePath + file.Name()
				if !isSupportedFormat(fileInDirPath) {
					continue
				}
				sourceList = append(sourceList, source{name: file.Name(), path: fileInDirPath})
			}
		} else {
			if !isSupportedFormat(filePath) {
				continue
			}
			sourceList = append(sourceList, source{name: filePath, path: filePath})
		}
	}
	return sourceList
}

func isSupportedFormat(fileName string) bool {
	for _, format := range supportedFormats {
		if strings.HasSuffix(strings.ToLower(fileName), format) {
			return true
		}
	}
	return false
}

func process(fileName string) {
	defer wg.Done()

	hasExifOrientation, angle := getExifOrientation(fileName)
	if hasExifOrientation {
		rotateAndDropExif(fileName, angle)
	}

	resizeIfNeeded(fileName)

	log.Println(fileName, "Done")
}

func getExifOrientation(fileName string) (hasExifOrientation bool, angle int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	exifInfo, err := exif.Decode(file)
	if err != nil {
		return
	}
	orientationTag, err := exifInfo.Get("Orientation")
	if err != nil {
		log.Printf("EXIF get orientation error: %v\n", err)
		return
	}
	orientationValue, err := orientationTag.Int(0)
	if err != nil {
		log.Printf("EXIF get orientation value error: %v\n", err)
		return
	}
	hasExifOrientation = true
	switch orientationValue {
	default:
		fallthrough
	case 0:
		return false, 0
	case 3:
		angle = 180
	case 6:
		angle = 90
	case 8:
		angle = 270
	}
	return
}

func rotateAndDropExif(fileName string, angle int) {
	srcImage, err := imaging.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	dstImage := imaging.Rotate(srcImage, float64(-angle), color.Black)
	saveErr := imaging.Save(dstImage, fileName)
	if saveErr != nil {
		log.Fatalf("Save failed: %v", saveErr)
	}
}

func getSize(fileName string) int64 {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.Size()
}

func resizeIfNeeded(fileName string) {
	var size = getSize(fileName)
	if size <= maxFileSize {
		return
	}
	srcImage, err := imaging.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var bounds = srcImage.Bounds()
	var maxBound = bounds.Dx()
	if bounds.Dy() > bounds.Dx() {
		maxBound = bounds.Dy()
	}
	var lowerBound = getLowerBound(maxBound)
	if maxBound < lowerBound {
		log.Printf("The image %v is too small for resize %v x %v", maxBound, bounds.Dx(), bounds.Dy())
		return
	}

	var scale = float64(lowerBound) / float64(maxBound)
	var width = int(math.Ceil(float64(bounds.Dx()) * scale))
	var height = int(math.Ceil(float64(bounds.Dy()) * scale))

	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)
	saveErr := imaging.Save(dstImage, fileName)
	if saveErr != nil {
		log.Fatalf("Save failed: %v", saveErr)
	}
	log.Println(fileName, "Resized", bounds.Dx(), "x", bounds.Dy(), "->", width, "x", height)

	resizeIfNeeded(fileName)
}

func getLowerBound(value int) int {
	var bound int
	for _, bound = range imageBoundSteps {
		if bound < value {
			return bound
		}
	}
	return bound
}
