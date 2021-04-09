package link

import (
	"fmt"
	"log"
)

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

//示例 1：
//输入：head = [1,3,2]
//输出：[2,3,1]

//法一： 顺序遍历，将元素添加至slice中，再将sli反向输出；
func reversePrint1(head *ListNode) []int {
	dummy := head
	var sli []int
	for dummy != nil {
		sli = append(sli, dummy.Val)
		dummy = dummy.Next

	}

	if len(sli) <= 1 {
		return sli
	}
	n := len(sli)
	for i := 0; i < n/2+n%2; i++ { // 注意此处需要判别奇偶数
		sli[i], sli[n-1-i] = sli[n-1-i], sli[i]
	}
	return sli
}

//*************有问题**********
//法二：优化法一,使用defer打印
func reversePrint2(head *ListNode) (sli []int) {
	//sli := []int{}
	dummy := head
	for dummy != nil {
		log.Println("-----: ", dummy.Val)
		defer func() {
			sli = append(sli, dummy.Val)
		}()
		dummy = dummy.Next
	}
	return sli
}

//*************有问题**********

//法二：优化法一,使用defer打印
func reversePrint3(head *ListNode) (res []int) {
	for head != nil {
		temp := head
		defer func() {
			res = append(res, temp.Val)
		}()
		head = head.Next
	}
	return
}

//法三：将链表反转再输出

// 法四：使用递归
func reversePrint(head *ListNode) {
	dummy := head
	printListReverse(dummy)
}

func printListReverse(node *ListNode) {
	if node != nil {
		if node.Next != nil {
			printListReverse(node.Next)
		}
		fmt.Println("val: ", node.Val)
	}
}
