package list

import (
	"fmt"
)

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
func (l *List) AddHead(v interface{}) {
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

//AddTail 添加数据到链表表尾 Append操作
func (l *List) AddTail(v interface{}) {
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

//Insert 插入数据到链表某个节点位置
func (l *List) Insert(index int, v interface{}) {
	if l.Head == nil {
		l.Head.Data = v
		l.Length++
		return
	}
	if index > l.Len() {
		l.AddTail(v)
		return
	}
	if index < 0 {
		l.AddHead(v)
		return
	}
	current := l.Head
	// pre := l.Head
	i := 0
	for current != nil {
		if index == i {
			// fmt.Println(i)
			current = &Node{Data: v, Next: current}
			break
		}
		i++
		current = current.Next
	}
}

//Get 获取某节点数据
func (l *List) Get(index uint) *Node {

	return l.Head
}

//Del 删除节点
func (l *List) Del(index int) {

}

//Reverse 链表反转
func (l *List) Reverse() {

}

//Clone 深拷贝链表
func (l *List) Clone() *List {
	l2 := new(List)
	if l.Head == nil {
		return l2
	}
	current := l.Head
	for current != nil {
		node := &Node{Data: current.Data, Next: current.Next}
		node.Next = l2.Head //往后移
		l2.Head = node
		current = current.Next
		l2.Length++
	}
	return l2
}

//Len 链表长度int
func (l *List) Len() int {
	return int(l.Length)
}

//IsEmpty 链表是否为空
func (l *List) IsEmpty() bool {
	return l.Head == nil
}

//ListShow 打印链表
func (l *List) Show() {
	if l.Head == nil {
		fmt.Printf("List is empty\n")
		return
	}
	fmt.Println("===start===")
	current := l.Head
	for current != nil {
		fmt.Printf("&Current:%p, &Data:%p, Data:%v\n", &current, &current.Data, current.Data)
		current = current.Next
	}
	fmt.Println("====end====")
}
