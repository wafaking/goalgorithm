package main

import (
	"fmt"
	"sync"
	"testing"
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
