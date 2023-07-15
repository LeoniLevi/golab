package main

import (
	"fmt"
	//"math"
	"math/rand"
	"time"

	//"sliceapp/sort"
	"github.com/LeoniLevi/golab/sort"
)

func main() {
	slice_action()
}

func slice_action() {

	fmt.Println(" === slice_action func started...")

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
		sort.MyQuickSortOld(arr, 0, arrLen)
		fmt.Println(" ~~~ Quick-Sorted<Old> Slice:\n", arr)
	}

	{
		arr := copySlice(arr0)
		sort.MyQuickSort01(arr)
		fmt.Println(" ~~~ Quick-Sorted<01> Slice:\n", arr)
	}

	{
		arr := copySlice(arr0)
		sort.MyQuickSortOld(arr, 0, arrLen)
		fmt.Println(" ~~~ Quick-Sorted<1> Slice:\n", arr)
	}

	arr1 := copySlice(arr0)
	sort.MyBubbleSort(arr1)
	fmt.Println(" ~~~ Bubble-Sorted Slice:")
	fmt.Println(arr1)

	idx := sort.SearchBinary(val, arr1, 0, arrLen)
	fmt.Printf("~~~ searchBinary: <Value=%d>: FoundIndex=%d\n", val, idx)

	idx = sort.SearchBinary01(arr1, val)
	fmt.Printf("~~~ searchBinary01: <Value=%d>: FoundIndex=%d\n", val, idx)

	fmt.Println(" === slice_action func Completed!")
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
