package list

import "fmt"

//Node 单链表节点
type Node struct {
	Data interface{}
	Next *Node
}

//List 单链表 包含长度
type List struct {
	Length uint
	Head   *Node
}

//listCreate 分配链表
func ListCreate() *List {
	list := new(List)
	list.Head = new(Node)
	return list
}

//listRelease 置空整个链表中的数据
func (l *List) ReleaseData() {
	current := l.Head
	for {
		current.Data = nil
		if current.Next != nil {
			current = current.Next
			continue
		}
		break
	}
}

//ListAddNodeHead 加到链表的表头
func (l *List) AddNodeHead(v interface{}) {
	node := &Node{Data: v}
	if l.Length == 0 {
		l.Head = node
		l.Length++
		return
	}
	node.Next = l.Head //往后移
	l.Head = node      //新表头
	l.Length++
}

//ListAddNodeTail 添加到链表表尾
func (l *List) AddNodeTail(v interface{}) {
	node := &Node{Data: v}
	if l.Length == 0 {
		l.Head.Next = node
		l.Length++
		return
	}
	current := l.Head
	for {
		if current.Next != nil {
			current = current.Next
			continue
		}
		current.Next = node
		break
	}
	l.Length++
}

//Insert 插入链表中间位置
func (l *List) Insert() {}

//Copy 深拷贝链表
// func (l *List) Clone() *List {
// 	l2 := new(List)
// 	if l.Head == nil {
// 		return l2
// 	}
// 	l2.Length = l.Length
// 	for l.Head != nil {
// 		l2.Head = &Node{
// 			Data: l.Head.Data,
// 			// Next: l.Head.Next,
// 		}
// 		l.Head = l.Head.Next
// 	}

// 	return l2
// }

//ListShow 打印链表
func (l *List) Show() {
	if l.Head == nil {
		fmt.Printf("List is empty\n")
		return
	}
	current := l.Head
	for {
		fmt.Printf("%p %v\n", &current.Data, current.Data)
		if current.Next != nil {
			current = current.Next
			continue
		}
		break
	}
}
