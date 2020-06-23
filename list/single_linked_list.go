package list

/*
	线性表--链式存储结构的实现与基本操作
*/

type SqSingleListOPeration interface {
	InitList() *SingleLink     // 初始化
	ListEmpty() bool           // 判断是否为空,true/false
	ClearList()                // 清空
	GetElem(i int) *Node       // 获取元素
	LocateElem(Node) int       // 查找元素位置
	ListInsert(int, Node) bool // 插入
	Remove(int) (bool, *Node)  // 删除
	Length() int               // 长度
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

func (sq *SingleLink) GetElem(i int) *Node {
	if sq.Size <= i || i <= 0 {
		return nil
	}
	var (
		num  int
		node = sq.Data
	)
	for {
		if num == i-1 {
			break
		}
		node = node.Next
		num++
	}
	return node
}

// func (sq *SingleLink) Remove(i int) (bool, *Node) {
// 	if sq.Size <= i || i <= 0 {
// 		return false, nil
// 	}

// 	var (
// 		num  int
// 		node *Node
// 	)
// 	for num := 0; num < i; num++ {

// 		if num == i-1 {
// 			// node =
// 			break
// 		}
// 		num++
// 	}
// 	sq.Size--

// 	return true, nil
// }
