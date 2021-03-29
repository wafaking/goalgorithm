package link

import (
	"fmt"
	"log"
)

var head *ListNode

// ListNode 结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func SetListNode(sli []int) {
	head = makeListNode(sli)
}

func makeListNode(sli []int) *ListNode {

	dummy := &ListNode{-1, nil}
	cur := dummy
	for _, v := range sli {
		node := &ListNode{v, nil}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func PrintHead() {
	var valSli []int
	//fmt.Println("head11: ", head)
	cur := head
	for cur != nil {
		valSli = append(valSli, cur.Val)
		cur = cur.Next
	}
	fmt.Println("***********************************")
	log.Printf("*******NodeList: %#v ********\n", valSli)
	fmt.Println("***********************************")
}
