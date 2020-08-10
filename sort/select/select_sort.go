package sort

import "fmt"

func SelectionSort(sli []int) {
	fmt.Println("before selection sort: ", sli)
	length := len(sli)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if sli[min] > sli[j] {
				min = j
			}
		}
		if min != i {
			sli[i], sli[min] = sli[min], sli[i]
		}
	}

	fmt.Println("after selection sort: ", sli)
}
