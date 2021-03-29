package link

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

//示例 1：
//输入：head = [1,3,2]
//输出：[2,3,1]

func reversePrint(head *ListNode) []int {
	//法一： 顺序遍历，将元素添加至slice中，再将sli反向输出；
	// 法二：将链表反转再输出
	dummy := head
	var sli []int
	for dummy != nil {
		sli = append(sli, dummy.Val)
		dummy = dummy.Next
	}

	n := len(sli)
	for i := 0; i < n-1; i++ {
		sli[i], sli[n-1] = sli[n-1], sli[i]
		n--
	}

	return sli
}
