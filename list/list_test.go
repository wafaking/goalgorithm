package list

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	//sli := []int{5, 3, 7, 2, 0, 4, 1, 2}
	//sli := []int{4, 7, 5, 9, 2, 3}
	//sli := []int{2, 3, 1, 0, 2, 5, 3}
	sli := []int{2, 7, 11, 15}
	InitNums(sli)
	PrintNums()
	t.Run()
}

func TestFindKthLargest(t *testing.T) {
	var k = 2
	fmt.Println("res: ", findKthLargest(nums, k))
}

func TestTwoSum(t *testing.T) {
	sum := 9
	res := twoSum(nums, sum)
	log.Println("res: ", res)
}

func TestFindRepeatNumber(t *testing.T) {
	res := findRepeatNumber(nums)
	log.Println("res: ", res)
}
