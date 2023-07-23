package sort

import "fmt"

func quickSort00(arr []int, idxBegin, idxEnd int) {
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

	quickSort00(arr, idxBegin, gtIdx-1)
	quickSort00(arr, gtIdx, idxEnd)
}

func quickSort01(arr []int) {
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

		quickSort01(arr[0 : gtIdx-1])
		quickSort01(arr[gtIdx:len])
	}
}

func quickSortOld(arr []int, idxBegin, idxEnd int) {
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

	quickSortOld(arr, idxBegin, splitIdx)
	quickSortOld(arr, splitIdx+1, idxEnd)
}
