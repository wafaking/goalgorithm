package sort

import (
	"testing"
)

var arraySli = [][]int{
	{10, 9, 7, 5, 6, 3},
	{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4},
	{1, 3, 2, 4, 5, 6, 7, 8, 9},
	{3, 9, 1, 4, 7},
	{432, 432432, 0, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4},
}

// var sli = []int{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4}

// var sli = []int{1, 3, 2, 4, 5, 6, 7, 8, 9}
// var sli []int

// func init() {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	r.Seed(time.Now().UnixNano())
// 	// rand.Seed(time.Now().UnixNano())

// 	for i := 0; i < 15; i++ {
// 		num := rand.Intn(1000)
// 		sli = append(sli, num)
// 	}
// }

func TestSelectionSort(t *testing.T) {
	for _, sli := range arraySli {
		SelectionSort(sli)
	}
}

func TestInsertSort(t *testing.T) {
	for _, sli := range arraySli {
		InsertSort(sli)
	}
}

func TestShellSort(t *testing.T) {
	for _, sli := range arraySli {
		ShellSort(sli)
	}
}
