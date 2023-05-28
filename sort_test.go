package sort

import (
	"math"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		init []int
		want []int
	}{
		{
			init: []int{2, 4, 7, 3, 5, 9},
			want: []int{2, 3, 4, 5, 7, 9},
		},
		{
			init: []int{1, 3, 12, 7, 15, 29, 45},
			want: []int{1, 3, 7, 12, 15, 29, 45},
		},
		{
			init: []int{3, 48, 87, 21, 22, 85, 94},
			want: []int{3, 21, 22, 48, 85, 87, 94},
		},
	}
	for _, item := range tests {
		before := make([]int, len(item.init))
		mid := int(math.Floor((float64(len(item.init)) / 2)))
		copy(before, item.init)
		Merge(item.init, mid)
		if !reflect.DeepEqual(item.init, item.want) {
			t.Errorf("Merge(%v,%d) returns %v, wants %v\n", before, mid, item.init, item.want)
		}
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		init []int
		want []int
	}{
		{
			init: []int{85, 21, 48, 94, 3, 22, 87, 86, 74, 15, 91, 49, 51, 11, 24, 50, 61, 41, 47, 34},
			want: []int{3, 11, 15, 21, 22, 24, 34, 41, 47, 48, 49, 50, 51, 61, 74, 85, 86, 87, 91, 94},
		},
		{
			init: []int{52, 79, 57, 64, 17, 38, 90, 63, 81, 85},
			want: []int{17, 38, 52, 57, 63, 64, 79, 81, 85, 90},
		},
	}
	for _, item := range tests {
		before := make([]int, len(item.init))
		copy(before, item.init)
		MergeSort(item.init)
		if !reflect.DeepEqual(item.init, item.want) {
			t.Errorf("MergeSort(%v) returns %v, wants %v\n", before, item.init, item.want)
		}
	}
}
