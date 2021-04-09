package list

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	//sli := []int{5, 3, 7, 2, 0, 4, 1, 2}
	//sli := []int{4, 7, 5, 9, 2, 3}
	sli := []int{2, 3, 1, 0, 2, 5, 3}
	//sli := []int{2, 7, 11, 15}
	//sli := []int{3, 2, 4}
	//sli := []int{3, 0, -2, -1, -1, -1, 1, 2}

	InitNums(sli)
	PrintNums()

	//Init2DNums()
	//Println2DNums()
	t.Run()
}

func TestFindKthLargest(t *testing.T) {
	var k = 2
	fmt.Println("res: ", findKthLargest(nums, k))
}

func TestTwoSum(t *testing.T) {
	sum := 6
	res := twoSum(nums, sum)
	log.Println("res: ", res)
}

func TestFindRepeatNumber(t *testing.T) {
	res := findRepeatNumber(nums)
	log.Println("res: ", res)
}

func TestThreeSum(t *testing.T) {
	res := threeSum(nums)
	log.Println("res: ", res)
}

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix = [][]int{{-5}}
	res := findNumberIn2DArray(matrix, -5)
	log.Println("res: ", res)
}

func TestMerge(t *testing.T) {
	type Te struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	sli := []Te{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
		},
		{
			nums1: []int{0},
			nums2: []int{1},
			m:     0,
			n:     1,
		},
		{
			nums1: []int{2, 0},
			nums2: []int{1},
			m:     1,
			n:     1,
		},
		{
			nums1: []int{4, 5, 6, 0, 0, 0},
			nums2: []int{1, 2, 3},
			m:     3,
			n:     3,
		},
	}

	//nums1 := []int{0}
	//nums2 := []int{1}
	//m, n := 0, 1

	//nums1 := []int{2, 0}
	//nums2 := []int{1}
	//m, n := 1, 1

	for _, v := range sli {
		log.Println("nums1: ", v.nums1)
		log.Println("nums2: ", v.nums2)
		merge(v.nums1, v.m, v.nums2, v.n)
		log.Println("res: ", v.nums1)
	}

}
