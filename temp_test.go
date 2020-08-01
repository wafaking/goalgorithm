package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	var name = "wafa"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		name = "nancy"
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("name：", name)
}

func TestFunc2(t *testing.T) {
	var sli = []string{"a", "b", "c"}
	var wg sync.WaitGroup
	for _, v := range sli {
		wg.Add(1)
		fmt.Println("v: ", v)
		time.Sleep(time.Second)
		go func() {
			fmt.Printf("name: %s\n", v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestDefer(t *testing.T) {
	// DeferFunc()
	var a int = 2

	defer fmt.Printf("defer1 a: %d \n", a) // 2
	a++
	defer fmt.Printf("defer2 a: %d \n", a) // 3
	a++
	defer func(a *int) {
		fmt.Printf("defer3 a: %v \n", *a) // 6
	}(&a)
	a++
	defer func() {
		fmt.Printf("defer4 a: %d \n", a) // 6
	}()
	a++
}

func TestDefer2(t *testing.T) {
	a()
}
func a() {
	i := 0
	defer fmt.Println(i) // 1
	i++
	return
}

func TestDefer3(t *testing.T) {
	// res := hehe()
	res1 := c()
	fmt.Println("res1: ", res1) // 2

	res2 := hehe()
	fmt.Println("res2: ", res2) // 2
}

func hehe() (i int) {
	fmt.Println("i: ", i) // i=0
	defer func() {
		i++ // i=2
	}()
	i = i + 2 // i=2
	return 1
}

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func TestNamedOrNotReturnFunc(t *testing.T) {
	val1 := NamedReturnFunc()
	fmt.Printf("var1: %d\n", val1) // 5

	val3 := NonNamedReturnFunc()
	fmt.Printf("var3: %d\n", val3) // 99
}

func NamedReturnFunc() (val1 int32) {
	defer func() {
		val1 = 5
	}()

	val1 = 99
	fmt.Printf("NamedReturnFunc return, val1: %d\n", val1) // 99
	return
}

func NonNamedReturnFunc() int {
	var val1 int
	defer func() {
		val1 = 5
	}()

	val1 = 99
	fmt.Printf("NonNamedReturnFunc return, val1: %d\n", val1) // 99
	return val1
}

func TestSumFunc(t *testing.T) {
	var sli = []int{1, 2, 7, 11, 15, 8}
	var total = 9
	// res := SumFunc1(sli, total)
	res := SumFunc2(sli, total)
	// res := SumFunc3(sli, total)
	fmt.Println("res: ", res)
}

func SumFunc1(nums []int, val int) []string {
	pair := []string{}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j <= length-1; j++ {
			if val == nums[i]+nums[j] {
				pair = append(pair, fmt.Sprintf("%d,%d", nums[i], nums[j]))
			}
		}
	}
	return pair
}

func SumFunc2(nums []int, val int) []string {
	pair := []string{}
	var myMap = make(map[int]int, 0)
	for k, v := range nums {
		myMap[v] = k
	}

	for _, v := range nums {
		res := val - v
		if _, ok := myMap[res]; ok {
			pair = append(pair, fmt.Sprintf("%d,%d", v, res))
		}
	}

	return pair
}

func SumFunc3(nums []int, val int) []string {
	pair := []string{} // 存放结果
	length := len(nums)
	mymap := make(map[int]int, length) // 存放差值
	for k, v := range nums {
		res := val - v
		if _, ok := mymap[res]; !ok {
			mymap[v] = k
		} else {
			pair = append(pair, fmt.Sprintf("%d,%d", res, v))
		}
	}
	return pair
}
