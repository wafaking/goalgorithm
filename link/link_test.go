package link

import (
	"log"
	"testing"
)

func TestMain(t *testing.M) {

	ssli := [][]int{
		{1, 2, 2, 3, 3, 4, 5},
		{8, 4, 0, 6, 5, 6, 5, 7},
	}
	for _, sli := range ssli {
		SetListNode(sli)
		PrintHead()
		t.Run()
	}

}

func TestPrintListNode(t *testing.T) {
	PrintHead()
}

func TestDeleteDuplicatesAll(t *testing.T) {
	deleteDuplicatesAll(head)
	PrintHead()
}
func TestReversePrint(t *testing.T) {
	res := reversePrint(head)
	log.Println("res: ", res)
}
