package algo

import (
	"reflect"
	"testing"
)

func TestSort(b *testing.T) {
	var in = []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(Sort(in), []int{1, 2, 3, 4, 5}) {
		b.Error("Got ", Sort(in))
	}
}
