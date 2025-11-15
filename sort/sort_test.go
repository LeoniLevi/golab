package sort_test

import	"github.com/LeoniLevi/golab/sort"
import "testing"

//import "slices"

func TestBubbleSort(t *testing.T) {
	arr0 := []int{4, 8, 3, 5, 1, 2}
	expected := []int{1, 2, 3, 4, 5, 8}
	sort.BubbleSortInt(arr0)
	for i := range arr0 {
		if arr0[i] != expected[i] {
			t.Errorf("BubbleSortInt failed in index %d", i)
		}
	}
}

func TestQuickSort(t *testing.T) {
	arr := [8]int{67, -4, 33, -22, 6, 9, -7, 11}
	expected := [8]int{-22, -7, -4, 6, 9, 11, 33, 68}
	sort.QuickSortInt(arr[:])
	if arr != expected {
		t.Errorf("QuickSortInt failed! result: %v, expected: %v", arr, expected)
	}
}
