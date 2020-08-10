package sort

import "fmt"

func InsertSort(sli []int) {
	fmt.Println("before insert sort: ", sli)

	length := len(sli)
	if length < 2 {
		return
	}

	for i := 1; i < length; i++ {
		if sli[i] < sli[i-1] {
			for j := i; j >= 0 && sli[j] < sli[j-1]; j-- {
				// if sli[j] >= sli[j-1] {
				// 	break
				// }
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}

	fmt.Println("after insert sort: ", sli)
}
