package tree

import (
	"fmt"
)

var root *TreeNode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 添加节点
func Insert(num int) {
	node := &TreeNode{Val: num}
	if root == nil {
		root = node
		return
	}
	// 原sli：4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	// pre 4 3 1 2 8 5 7 6 9 10
	// in 1 2 3 4 5 6 7 8 9 10
	// post 2 1 3 6 7 5 10 9 8 4

	cur := root
	for {
		if num < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return
			}
			cur = cur.Right
		}
	}
}

// 前序遍历(先根节点，再左子树，再右子树)
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 中序遍历(先左子树，再根节点，再右子树)
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

// 后序遍历（先左子树，再右子树，最后根节点）
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

// 层序遍历
func LevelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var sli []*TreeNode
	sli = append(sli, root)
	for len(sli) > 0 {
		first := sli[0]
		//fmt.Println(first.Val)
		res = append(res, first.Val)
		if first.Left != nil {
			sli = append(sli, first.Left)
		}
		if first.Right != nil {
			sli = append(sli, first.Right)
		}
		sli = sli[1:]
	}
	return res
}
