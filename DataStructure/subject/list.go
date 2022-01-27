package subject

import (
	List "GoPractise/DataStructure/list"
)

func TestList() {
	list := List.ListCreate()
	list.AddHead(map[string]int{"a": 1})
	list.AddTail(map[string]int{"b": 2})
	list.AddTail(map[string]int{"b": 3})
	list.AddHead(map[string]int{"d": 4})
	list.Show()

	newList := list.Clone()
	newList.Show()

	list.ReleaseData()
	list.Show()
	/*
		$ go run main.go
			===start===
			0xc0000ae0c0 map[d:4]
			0xc0000ae078 map[a:1]
			0xc0000ae090 map[b:2]
			0xc0000ae0a8 map[b:3]
			====end====
			===start===
			0xc0000ae1e0 map[b:3]
			0xc0000ae1c8 map[b:2]
			0xc0000ae1b0 map[a:1]
			0xc0000ae198 map[d:4]
			====end====
			===start===
			0xc0000ae0c0 <nil>
			0xc0000ae078 <nil>
			0xc0000ae090 <nil>
			0xc0000ae0a8 <nil>
			====end====
	*/

	newList.Insert(1, &List.Node{})
	newList.Show()
}
