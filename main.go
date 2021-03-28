package main

import "log"

func main() {
	// 值类型
	var a = 5
	b := a

	// 指针类型
	var m = []int{1, 2, 3}
	n := m
	log.Printf("%v, %v\n", &a, &b) // 0xc000098000, 0xc000098008
	log.Printf("%v, %v\n", &m, &n) // &[1 2 3], &[1 2 3]
	n[2] = 0
	log.Printf("%v, %v\n", &m, &n) // &[1 2 0], &[1 2 0]
}
