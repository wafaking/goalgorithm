package sort

import (
	"testing"
)

var sli = []int{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4}

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

func TestBubbleSort(t *testing.T) {
	BubbleSort(sli)
}

func TestBubbleSortStandard(t *testing.T) {
	BubbleSortStandard(sli)
}

func TestBubbleSortStandardOptimize(t *testing.T) {
	BubbleSortStandardOptimize(sli)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(sli)
}

func TestInsertSort(t *testing.T) {
	InsertSort(sli)
}

func TestShellSort(t *testing.T) {
	ShellSort(sli)
}
