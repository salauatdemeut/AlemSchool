package main

import (
	"fmt"

)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	for l.Head == nil {
		l.Head = n
		l.Tail = n
		}
		  l.Tail.Next = n
		  l.Tail = l.Tail.Next
}


func main() {

	link := &List{}

	ListPushBack(link, "Hello")
 ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}