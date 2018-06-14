package main

// Sort []int for O(N^2)
func Sort(source []int) []int {
	var count = len(source)
	for y := 1; y < count; y++ {
		for x := 0; x < count-y; x++ {
			if source[x] > source[x+1] {
				source[x], source[x+1] = source[x+1], source[x]
			}
		}
	}
	return source
}
