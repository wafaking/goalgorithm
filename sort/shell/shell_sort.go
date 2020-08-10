package sort

import "fmt"

func ShellSort(sli []int) {
	fmt.Println("before shell sort: ", sli)
	length := len(sli)
	if length < 2 {
		return
	}
	step := length / 2

	for step >= 1 {
		insertSort(sli, step)
		step = step / 2
	}

	fmt.Println("after shell sort: ", sli)
}

func insertSort(sli []int, step int) {
	//   0 1 2  3 4 5  6 7
	//  [1 3 46 5 2 10 9 4  ] length=8
	//  [0, 2] [3, 10], [46, 9], [5, 4] step=4
	//  [1, 46, 2, 9] [3, 5, 10, 4] step=2
	//  [1 3 46 5 2 10 9 4  ] step=1
	fmt.Println("before shell sort1: ", sli)

	length := len(sli)

	for i := 1; i < length; i += step {
		if sli[i] < sli[i-1] {
			for j := i; j < length && sli[j] < sli[j-1]; j -= step {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}

		}
	}

	fmt.Println("after shell sort1: ", sli)

}

func insertSort1(tree []int, step int) {
	// step=4
	//  [0, 2] [3, 10], [46, 9], [5, 4] step=4
	//  [1 3 46 5 2 10 9 4  ] length=8
	//  [1, 46, 2, 9] [3, 5, 10, 4] step=2
	for i := step; i < len(tree); i++ {
		for j := i; j >= step; j -= step {
			fmt.Println("----", i, j, step, j-step, "------", tree[j], tree[j-step])
			if tree[j] >= tree[j-step] {
				break
			}
			tree[j], tree[j-step] = tree[j-step], tree[j]
			fmt.Println("sli: ", tree)
		}
	}
}
