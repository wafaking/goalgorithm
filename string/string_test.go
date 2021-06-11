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

func TestMyAtoi(t *testing.T) {
	s := map[string]int{
		"42":              42,
		"   -42":          -42,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
		"21474836460":     2147483647,
		"2147483648":      2147483647,
	}
	var res int
	for k, v := range s {
		res = myAtoi(k)
		if res != v {
			t.Fatalf("k: %s, res: %d, expect: %d\n", k, res, v)
		}
	}
}
func TestPermutation(t *testing.T) {
	res := permutation4("abc")
	log.Println("res: ", res)
}

func TestCheckInclusion(t *testing.T) {
	res := checkInclusion("ab", "eidbaooo")
	log.Println("res: ", res)
}
