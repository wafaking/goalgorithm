package tree

import (
	"log"
	"testing"
)

func TestPermute(t *testing.T) {
	log.Println(permute([]int{1, 2, 3}))
}

func TestPermuteUnique(t *testing.T) {
	var list = []int{1, 2, 1, 3}
	log.Println(PermuteUnique(list))
}
