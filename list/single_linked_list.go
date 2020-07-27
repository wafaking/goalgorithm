package list

/*
	线性表--链式存储结构的实现与基本操作
*/

type SqSingleListOPeration interface {
	InitList() *SingleLink      // 初始化
	ListEmpty() bool            // 判断是否为空,true/false
	ClearList()                 // 清空
	GetElem(i int) *Node        // 获取元素
	LocateElem(Node) int        // 查找元素位置
	ListInsert(int, *Node) bool // 插入
	Remove(int) (bool, *Node)   // 删除
	Length() int                // 长度
}

// Node 单结构
type Node struct {
	Next *Node
	Data string
}

type SingleLink struct {
	Size int
	Data *Node
}

func (sq *SingleLink) InitList() *SingleLink {
	return &SingleLink{}
}

func (sq *SingleLink) Length() (length int) {
	if sq == nil {
		return
	}
	return sq.Size
}

// TODO 参考书上
func (sq *SingleLink) GetElem(i int) *Node {
	if sq.Size <= i || i <= 0 {
		return nil
	}
	var (
		num  int
		node *Node
	)

	for {
		if num == i {
			break
		}
		if num == 0 {
			node = sq.Data
		} else {
			node = node.Next
		}
		num++
	}
	return node
}

func (sq *SingleLink) ClearList() {
	if sq.Size == 0 {
		return
	}
	sq.Size = 0
	sq.Data = nil
	return
}

func (sq *SingleLink) ListEmpty() bool {
	return sq.Size == 0
}

func (sq *SingleLink) LocateElem(node Node) int {
	var index = -1
	if sq.Size == 0 {
		return index
	}

	var inode = sq.Data
	for num := 0; num < sq.Size; num++ {
		num++
		if node != *inode {
			inode = inode.Next
			continue
		}
		index = num
	}
	return index
}

func (sq *SingleLink) Remove(i int) (bool, *Node) {
	var (
		hasRomoved bool
		rNode      *Node
	)
	if sq.Size <= i || i <= 0 {
		return hasRomoved, rNode
	}

	var inode = sq.Data

	for num := 0; num < i; num++ {
		num++
		curNode := inode
		if num != i {
			inode = inode.Next
			continue
		}
		curNode.Next = inode.Next
		rNode = inode
		hasRomoved = true
		sq.Size--
	}
	return hasRomoved, rNode
}

func (sq *SingleLink) ListInsert(pos int, node *Node) bool {
	var isInsert bool

	if pos > sq.Size+1 || pos <= 0 || node == nil {
		return isInsert
	}
	// if pos == 0 {
	// 	node.Next = sq.Data
	// 	isInsert = true
	// 	return isInsert
	// }
	if pos == sq.Size+1 {
		lastNode := sq.GetElem(sq.Size)
		lastNode.Next = node
		isInsert = true
		return isInsert
	}

	var iNode = sq.Data
	for i := 0; i < sq.Size; i++ {
		i++
		curNode := iNode
		if i != pos {
			iNode = iNode.Next
			continue
		}
		node.Next = iNode
		curNode.Next = node
		sq.Size++
		isInsert = true
	}

	return isInsert
}
