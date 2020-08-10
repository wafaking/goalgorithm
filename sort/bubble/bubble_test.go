package bubble

import (
	"fmt"
	"testing"
)

var arraySli = [][]int{
	{10, 9, 7, 5, 6, 3},
	{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4},
	{1, 3, 2, 4, 5, 6, 7, 8, 9},
	{3, 9, 1, 4, 7},
	{432, 432432, 0, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4},
}

func TestBubbleSort(t *testing.T) {
	for _, sli := range arraySli {
		BubbleSort(sli)
	}
}

func TestStandardBubbleSort(t *testing.T) {
	for _, sli := range arraySli {
		BubbleSortStandard(sli)
	}
}

func TestStandardBubbleSortOptimize(t *testing.T) {
	for _, sli := range arraySli {
		BubbleSortStandardOptimize(sli)
	}
}

func TestBubbleSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)

	sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }
