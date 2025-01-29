package main

import (
	"fmt"
)

// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法∏
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {

	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

	// 调用f函数本体
	f(p)
}

// // 声明接口变量
// var invoker Invoker

// // 实例化结构体
// s := new(Struct)

// // 将实例化的结构体赋值到接口
// invoker = s

// // 使用接口调用实例化结构体的方法Struct.Call
// invoker.Call("hello")

// // 将匿名函数转为FuncCaller类型，再赋值给接口
// invoker = FuncCaller(func(v interface{}) {
// 	fmt.Println("from function", v)
// })

// // 使用接口调用FuncCaller.Call，内部会调用函数本体
// invoker.Call("hello")

// Invoker是什么
// Invoker是一个接口，它包含一个方法Call
// Call方法接收一个interface{}类型的参数
// interface{}是一个空接口，可以接收任意类型的参数
// 因此，Invoker接口可以接收任意类型的参数
// Invoker接口是一个多态接口，它可以是任意类型的参数
// Invoker接口是一个抽象类型，它不能直接用来实例化对象
// Invoker接口是一个定义，它描述了对象的行为
// Invoker接口是一组抽象方法的集合
// Invoker接口是一种类型，它规定了实现该接口的类必须提供哪些方法
// Invoker接口是一种契约，实现了该接口的类必须遵守契约
// Invoker接口是一种数据类型，它包含一组方法
// Invoker接口是一种数据结构，它包含多个方法

// type Invoker interface {
// 	Call(interface{})
// }

//Struct是一个结构体类型，它包含一个方法Call
//Call方法接收一个interface{}类型的参数
//interface{}是一个空接口，可以接收任意类型的参数
//因此，Call方法可以接收任意类型的参数
//Call方法是一个多态方法，它可以是任意类型的参数
//Call方法是一个具体方法，它是一个实例方法
//Call方法是一个函数，它是一个可调用对象
//Call方法是一个方法，它是一个函数
//Call方法是一个函数，它是一个方法
//Call方法是一个方法，它是一个成员方法
//Call方法是一个成员方法，它是一个实例方法
//Call方法是一个成员方法，它是一个具体方法
//Call方法是一个成员方法，它是一个多态方法
//Call方法是一个成员方法，它是一个函数
//Call方法是一个成员方法，它是一个可调用对象
//Call方法是一个成员方法，它是一个方法
//Call方法是一个成员方法，它是一个抽象方法
//Call方法是一个成员方法，它是一个数据结构
//Call方法是一个成员方法，它是一个类型
//Call方法是一个成员方法，它是一个契约
//Call方法是一个成员方法，它是一组抽象

func main() {
	// da(dog{})
	// c1 := cat{}
	// da(c1)
	// c2 := dog{}
	// da(&c2)
	// p1 := person{
	// 	name: "zhangsan",
	// 	age:  18,
	// }
	// da(p1)

	var gj sayer
	c2 := cat{}
	gj = c2
	fmt.Println(gj)
	p2 := person{
		name: "zhangsan",
		age:  18,
	}
	gj = &p2
	fmt.Println(gj)

	//这个地方取地址，是因为c2是值类型，而&c2是引用类型
}

type dog struct{}

func (d dog) say() {
	fmt.Println("wang!")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("miaomiao!")
}

type person struct {
	name string
	age  int
}

func (p person) say() {
	fmt.Printf("别打我，我不是%s，我进年才%d岁\n", p.name, p.age)
}

type sayer interface {
	say()
}

func da(daozhi sayer) {
	daozhi.say()
}
