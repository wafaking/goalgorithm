package link

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 法一：
//func copyRandomList(head *Node) *Node {
//	return nil
//}
func copyRandomListM(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next
	pre := &newHead
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		p = p.Next
		pre = pre.Next
	}
	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			step := findStep(head, p.Random)
			newP.Random = target(&newHead, step)
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}

//确定从头结点到目标节点所经过的节点数
func findStep(head, target *Node) int {
	p := head
	step := 0
	for p != target {
		p = p.Next
		step++
	}
	return step
}

//返回从头结点开始，走step步所到达的节点
func target(head *Node, step int) *Node {
	p := head
	for step > 0 {
		p = p.Next
		step--
	}
	return p
}

func copyRandomListN(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next
	pre := &newHead
	connection := make(map[*Node]*Node)
	connection[head] = pre

	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		connection[p] = newNode
		p = p.Next
		pre = pre.Next
	}

	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			newP.Random = connection[p.Random]
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}

// 法二：
func copyRandomList(head *Node) *Node {
	return nil
}

// 法三：在每个节点后添加一个新节点，再复制随机节点，最后再拆分奇偶节点
func copyRandomList3(head *Node) *Node {
	if head == nil {
		return head
	}
	// 1. 复制节点
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = cur.Next.Next
	}

	//2. 复制Random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 拆分奇、偶节点
	newNode := head.Next // 新的链表
	oldCur := head       // cur和newCur都是用于遍历的当前节点，用newNode自己处理后，返回的是最后节点的val
	newCur := newNode

	for newCur.Next != nil {
		oldCur.Next = oldCur.Next.Next
		newCur.Next = newCur.Next.Next

		//改变当前节点
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	// 最后一次遍历，当newCur.Next=nil时，oldCur.Next仍旧指向newCur，
	// 因此需要让oldCur.Next指向nil
	oldCur.Next = nil

	return newNode
}
