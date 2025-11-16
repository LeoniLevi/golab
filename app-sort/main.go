package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LeoniLevi/golab/sort"
)

func main() {
	//arr := [8]int{67, -4, 33, -22, 6, 9, -7, 11}
	//expected := [8]int{-22, -7, -4, 6, 9, 11, 33, 68}
	//sort.QuickSortInt(arr[:])
	//Println

	examineSorting()
}
func examineSorting() {
	fmt.Println(" === examineSorting started...")

	arrLen := 30

	maxValue := 90
	arr0 := buildRandomSlice(arrLen, maxValue)

	//arr0 = []int{4, 8, 3, 5, 1, 2}
	fmt.Println(" ~~~ Initial Slice:")
	fmt.Println(arr0)

	val := rand.Intn(maxValue)
	setValueSomewhere(val, arr0)
	fmt.Println(" ~~~ added value", val, ":\n", arr0)
	
	arr := copySlice(arr0)
	sort.QuickSortInt(arr)
	fmt.Println(" ~~~ examineSorting/QuickSortInt - Slice:\n", arr)

	arr1 := copySlice(arr0)
	sort.BubbleSortInt(arr1)
	fmt.Println(" ~~~ examineSorting/BubbleSortInt - Slice:")
	fmt.Println(arr1)

	idx := sort.SearchBinary(arr1, val)
	fmt.Printf("~~~ examineSorting/SearchBinary: <Value=%d>: FoundIndex=%d\n", val, idx)

	fmt.Println(" === examineSorting func Completed!")
}

func buildRandomSlice(numElems, maxElemVal int) (newSlice []int) {
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
