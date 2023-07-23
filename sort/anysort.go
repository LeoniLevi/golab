package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func ExamineSorting() {
	fmt.Println(" === ExamineSorting started...")

	arrLen := 30

	maxValue := 90
	arr0 := buildSlice(arrLen, maxValue)

	//arr0 = []int{4, 8, 3, 5, 1, 2}
	fmt.Println(" ~~~ Initial Slice:")
	fmt.Println(arr0)

	val := rand.Intn(maxValue)
	setValueSomewhere(val, arr0)
	fmt.Println(" ~~~ added value", val, ":\n", arr0)

	{
		arr := copySlice(arr0)
		quickSortOld(arr, 0, arrLen)
		fmt.Println(" ~~~ ExamineSorting/QuickSort<Old> - Slice:\n", arr)
	}

	{
		arr := copySlice(arr0)
		quickSort00(arr, 0, arrLen)
		fmt.Println(" ~~~ ExamineSorting/QuickSort<00> - Slice:\n", arr)
	}
	{
		arr := copySlice(arr0)
		quickSort01(arr)
		fmt.Println(" ~~~ ExamineSorting/QuickSort<01> - Slice:\n", arr)
	}

	arr1 := copySlice(arr0)
	bubbleSort(arr1)
	fmt.Println(" ~~~ ExamineSorting/BubbleSort - Slice:")
	fmt.Println(arr1)

	idx := searchBinary(val, arr1, 0, arrLen)
	fmt.Printf("~~~ ExamineSorting/searchBinary: <Value=%d>: FoundIndex=%d\n", val, idx)

	idx = searchBinary01(arr1, val)
	fmt.Printf("~~~ ExamineSorting/searchBinary: <Value=%d>: FoundIndex=%d\n", val, idx)

	fmt.Println(" === ExamineSorting func Completed!")
}

func buildSlice(numElems, maxElemVal int) (newSlice []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	newSlice = make([]int, numElems)
	for i := 0; i < numElems; i++ {
		//newSlice[i] = r1.Intn(maxElemVal)
		newSlice[i] = r1.Int()%maxElemVal + 1
	}
	return
}

func copySlice(arr []int) []int {
	sz := len(arr)
	sa := make([]int, sz)
	copy(sa, arr)
	return sa
}

func setValueSomewhere(value int, slice []int) {
	idx := rand.Intn(len(slice))
	slice[idx] = value
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

func searchBinary(value int, slice []int, idxBegin, idxEnd int) (foundIdx int) {
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
		foundIdx = searchBinary(value, slice, idxBegin, idx)
	} else {
		foundIdx = searchBinary(value, slice, idx+1, idxEnd)
	}
	return
}

func searchBinary01(arr []int, value int) (foundIdx int) {
	idx := len(arr) / 2
	switch {
	case arr[idx] == value:
		foundIdx = idx
	case idx == 0:
		foundIdx = -1 // not found
	case value < arr[idx]:
		foundIdx = searchBinary01(arr[:idx], value)
	default:
		foundIdx = idx + 1 + searchBinary01(arr[idx+1:], value)
	}
	return
}
