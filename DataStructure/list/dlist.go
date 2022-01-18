package list

//DoubleNode 双链表节点
type DoubleNode struct {
	Data interface{}
	Prev *DoubleNode
	Next *DoubleNode
}

//DoubleList 双链表
type DoubleList struct {
	Length uint
	Head   *DoubleNode
	Tail   *DoubleNode
}

//DoubleListCreate
func DoubleListCreate() *DoubleList {
	return &DoubleList{
		Head: nil, Tail: nil, Length: 0,
	}
}
