package tree

import (
	"log"
	"testing"
)

// 原sli：4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
// pre 4 3 1 2 8 5 7 6 9 10
// in 1 2 3 4 5 6 7 8 9 10
// post 2 1 3 6 7 5 10 9 8 4
// layer 4 3 8 1 5 9 2 7 10 6

func TestMain(t *testing.M) {
	list := []int{
		4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	}
	for _, v := range list {
		Insert(v)
	}
	t.Run()
}

func TestPermute(t *testing.T) {
	log.Println(permute([]int{1, 2, 3}))
}

func TestPermuteUnique(t *testing.T) {
	var list = []int{1, 2, 1, 3}
	log.Println(PermuteUnique(list))
}

func TestPreOrder(t *testing.T) {
	PreOrder(root)
}
func TestInOrder(t *testing.T) {
	InOrder(root)
}
func TestPostOrder(t *testing.T) {
	PostOrder(root)
}

func TestLevelOrder(t *testing.T) {
	res := LevelOrder(root)
	log.Println("res: ", res)
}
