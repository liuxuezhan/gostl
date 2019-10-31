package main

import (
	"fmt"
	list "github.com/liyue201/gostl/ds/list/bid_list"
	"github.com/liyue201/gostl/ds/list/simple_list"
)

func simpleList() {
	l := simple_list.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n===============\n", )
}

func bidList() {
	l := list.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n", )
}

func main() {
	simpleList()
	bidList()
}