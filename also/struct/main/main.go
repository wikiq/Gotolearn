package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FunCaller func(interface{})

func (f FunCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call("hello")
	invoker = FunCaller(func(v interface{}) {
		fmt.Println("from funcaller", v)
	})
	invoker.Call("hello")
}
