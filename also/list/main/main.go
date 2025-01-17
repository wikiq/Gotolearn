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

func list_delete() {
	l := list.New()
	l.PushBack(("canonical"))
	//canonical
	l.PushFront(67)                   //67,canonical
	element := l.PushBack("fist")     //67,canonical,fist
	l.InsertAfter("After", element)   //67,canonical,fist,After
	l.InsertBefore("before", element) //67,After,canonical,before,fist
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("====================================")
	l.Remove(element) //删除元素67,After,canonical,before,After
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func nil_test() {
	// var nil = errors.New("my god")
	// fmt.Println(nil)

	var arr []int
	var num *int
	fmt.Println(arr == nil)
	fmt.Println(num == nil)
	// fmt.Println(arr == nil && num == nil)
}

func main() {
	// list_test()
	// list_delete()
	nil_test()
}
