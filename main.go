package main

import "fmt"

func main() {
	var a = 10
	switch {
	case a > 5:
		fmt.Println("22")
	case a > 4:
		fmt.Println("11")
	case a > 10:
		fmt.Println("33")
	}

	var ch2 = make(chan int, 1)
	var ch1 = make(chan int, 1)
	ch2 <- 5
	ch1 <- 6

	select {
	case m := <-ch1:
		fmt.Println(m)
	case n := <-ch2:
		fmt.Println(n)
	default:
		fmt.Println("33")
	}
}
