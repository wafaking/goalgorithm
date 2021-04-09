package str

import (
	"fmt"
	"log"
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

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	res := replaceSpace(s)
	log.Println("res: ", res)
}
