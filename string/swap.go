package str

import "fmt"

func SwapToNum(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
