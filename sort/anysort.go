package sort

import (
	"fmt"
)

func BubbleSortInt(nums []int) {
	bubbleSort(nums)
}

func QuickSortInt(nums []int) {
	quickSort01(nums)
}

func SearchBinary(arr []int, value int) (foundIdx int) {
	idx := len(arr) / 2
	switch {
	case arr[idx] == value:
		foundIdx = idx
	case idx == 0:
		foundIdx = -1 // not found
	case value < arr[idx]:
		foundIdx = SearchBinary(arr[:idx], value)
	default:
		foundIdx = idx + 1 + SearchBinary(arr[idx+1:], value)
	}
	return
}


func bubbleSort(arr []int) {
	sz := len(arr)
	for k := sz - 1; k > 0; k-- {
		for i := 0; i < k; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func searchBinary00(value int, slice []int, idxBegin, idxEnd int) (foundIdx int) {
	//fmt.Printf(" *** searchBinary: %d .. %d\n", idxBegin, idxEnd)
	if idxBegin >= idxEnd {
		msg := fmt.Sprintf("SearchBinary - incorrect idx : %d .. %d", idxBegin, idxEnd)
		panic(msg)
	}
	idx := idxBegin + (idxEnd-idxBegin)/2
	if slice[idx] == value {
		foundIdx = idx
	} else if idx == idxBegin {
		foundIdx = -1 // value not found
	} else if value < slice[idx] {
		foundIdx = searchBinary00(value, slice, idxBegin, idx)
	} else {
		foundIdx = searchBinary00(value, slice, idx+1, idxEnd)
	}
	return
}
