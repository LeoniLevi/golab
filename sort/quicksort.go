package sort

import "fmt"

func hello1() {
	fmt.Println("-- H.eLLO1")
}

func myQuickSort00(arr []int, idxBegin, idxEnd int) {
	if idxEnd-idxBegin < 2 {
		return
	}
	axisVal := arr[idxBegin]
	gtIdx := idxBegin + 1
	for i := idxBegin + 1; i < idxEnd; i++ {
		if arr[i] <= axisVal {
			arr[gtIdx], arr[i] = arr[i], arr[gtIdx]
			gtIdx += 1
		}
	}
	arr[idxBegin], arr[gtIdx-1] = arr[gtIdx-1], arr[idxBegin]

	//fmt.Println(" ** myQuickSort0:", arr[idxBegin:gtIdx-1], arr[gtIdx-1], arr[gtIdx:idxEnd])

	myQuickSort00(arr, idxBegin, gtIdx-1)
	myQuickSort00(arr, gtIdx, idxEnd)
}

func MyQuickSort01(arr []int) {
	len := len(arr)
	if len > 2 {
		axisVal := arr[0]
		gtIdx := 1 // to keep the last index of greater-than-axis element

		for i := 1; i < len; i++ {
			if arr[i] < axisVal {
				arr[gtIdx], arr[i] = arr[i], arr[gtIdx]
				gtIdx++
			}
		}
		arr[0], arr[gtIdx-1] = arr[gtIdx-1], arr[0]

		MyQuickSort01(arr[0 : gtIdx-1])
		MyQuickSort01(arr[gtIdx:len])
	}
}

func myQuickSort1(arr []int, low, high int) {
	if high-low >= 1 {
		pivot := arr[high]
		gtIdx := low // to keep the last index of greater-than-pivot element

		for i := low; i < high; i++ {
			if arr[i] < pivot {
				arr[i], arr[gtIdx] = arr[gtIdx], arr[i]
				gtIdx += 1
			}
		}
		arr[gtIdx], arr[high] = arr[high], arr[gtIdx]

		myQuickSort1(arr, low, gtIdx-1)
		myQuickSort1(arr, gtIdx+1, high)
	}
}

func MyQuickSortOld(arr []int, idxBegin, idxEnd int) {
	if idxBegin > idxEnd {
		msg := fmt.Sprintf("myQuickSort - incorrect idx : %d .. %d", idxBegin, idxEnd)
		panic(msg)
	}
	if idxEnd-idxBegin <= 1 { // if slice empty or contains just one element
		return
	}

	axisVal := arr[idxBegin]
	gtIdx := -1
	for i := idxBegin + 1; i < idxEnd; i++ {
		if arr[i] <= axisVal {
			if gtIdx >= 0 {
				arr[gtIdx], arr[i] = arr[i], arr[gtIdx]
				gtIdx += 1
			}
		} else { // arr[i] > axisVal
			if gtIdx < 0 {
				gtIdx = i
			}
		}
	}

	/// Now move Axis value into 'split' position just before the first greater element

	var splitIdx int
	if gtIdx >= 0 {
		if gtIdx <= idxBegin {
			panic("myQuickSort ERROR: valid gtIdx <= beginIdx")
		}
		if gtIdx > idxBegin+1 {
			arr[idxBegin], arr[gtIdx-1] = arr[gtIdx-1], arr[idxBegin]
		}
		splitIdx = gtIdx - 1
	} else { // if all elements less than first - move Axis to the end
		arr[idxBegin], arr[idxEnd-1] = arr[idxEnd-1], arr[idxBegin]
		splitIdx = idxEnd - 1
	}

	//fmt.Println(" ** myQuickSort:", arr[idxBegin:splitIdx], arr[splitIdx], arr[splitIdx+1:idxEnd])

	MyQuickSortOld(arr, idxBegin, splitIdx)
	MyQuickSortOld(arr, splitIdx+1, idxEnd)
}
