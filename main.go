package main

import "fmt"

func main() {

	var sli = []int{4, 1, 2, 2, 5, 1, 4, 3}
	var num = sli[0]
	// 依次异或，得到结果num
	for i := 1; i < len(sli); i++ {
		num = num ^ sli[i]
	}
	fmt.Println("num", num)
	if num == 0 { // 结果不为0，则说明给定数组不正确
		fmt.Println("wrong given number", sli)
		return
	}

	seprator := 1
	for {
		if 0 == (num & seprator) { // 与运算，找到最后一位为1的位数
			seprator <<= 1
			continue
		}
		break
	}
	fmt.Println("seprator: ", seprator)

	// 将原数组分为两个数组，分别异或运算即可得到结果
	var sliA, sliB = []int{}, []int{}
	var A, B int
	for _, v := range sli {
		if 0 == (v & seprator) {
			sliB = append(sliB, v)
			B ^= v
		} else {
			sliA = append(sliA, v)
			A ^= v
		}
	}
	fmt.Println("sliA: ", sliA, A)
	fmt.Println("sliB: ", sliB, B)
}
