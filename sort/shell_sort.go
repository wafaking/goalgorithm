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

// func insertSort(sli []int, step int) {
// 	//   0 1 2  3 4 5  6 7
// 	//  [1 3 46 5 2 10 9 4  ] length=8
// 	//  [0, 2] [3, 10], [46, 9], [5, 4] step=4
// 	//  [1, 46, 2, 9] [3, 5, 10, 4] step=2
// 	//  [1 3 46 5 2 10 9 4  ] step=1

// 	// step = 2
// 	for i := step; i < len(sli); i += step {
// 		for j := i; j < len(sli); j += step {
// 			if sli[j] >= sli[j-step] {
// 				break
// 			}
// 			sli[i], sli[i-step] = sli[i-step], sli[i]
// 		}
// 	}
// }

func insertSort(tree []int, step int) {
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
