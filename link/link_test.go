package link

import (
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	sli := []int{1, 2, 2, 3, 3, 4, 5}
	SetListNode(sli)
	PrintHead()
	t.Run()
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
