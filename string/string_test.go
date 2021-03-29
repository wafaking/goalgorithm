package str

import (
	"fmt"
	"testing"
)

func TestSwapToNum(t *testing.T) {
	var a, b = 10, 20
	SwapToNum(a, b)
}

func TestReverseVowels(t *testing.T) {
	var s1 = "leetcode"
	// var s2 = "hello"
	res := reverseVowels(s1)
	fmt.Println("res: ", res)
}
