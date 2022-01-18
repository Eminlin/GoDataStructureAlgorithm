package subject

import (
	List "GoPractise/DataStructure/list"
)

func TestList() {
	list := List.ListCreate()
	list.AddNodeHead(map[string]int{"a": 1})
	list.AddNodeTail(map[string]int{"b": 2})
	list.Show()

	// newList := list.Clone()
	// newList.Show()

	list.ReleaseData()
	list.Show()
	/*
		$ go run list.go
			0xc000004090 map[a:1]
			0xc0000040a8 map[b:2]
			0xc000004120 map[a:1]
			0xc0000040a8 map[b:2]
			0xc000004090 <nil>
			0xc0000040a8 <nil>
	*/
}
