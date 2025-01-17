package main

import (
	"container/list"
	"fmt"
)

func list_test() {
	l := list.New()
	l.PushBack("list")
	l.PushFront(67)
	// fmt.Println(l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//给我解释一下这个循环的意思，e := l.Front()是啥意思，e != nil是啥意思，e = e.Next()是啥意思
	// e := l.Front()是取出链表的第一个元素，e != nil是判断e是否为空，e = e.Next()是取出下一个元素
	// 为什么要用这个循环，这个循环的作用是什么
}

func main() {
	list_test()
}
