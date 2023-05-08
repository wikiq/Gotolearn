package customer_type

import (
	"testing"
	"time"
)

type IntConv func(op int) int

// type的用法是什么？
// type是自定义类型的关键字
// type IntConv func(op int) int的意思是什么？
// IntConv是自定义类型，func(op int) int是IntConv的类型
// func(op int) int是什么意思？
// func(op int) int是一个函数，参数是int，返回值是int
// IntConv是什么意思？
// IntConv是一个函数类型，参数是int，返回值是int
// IntConv的格式是什么？
// type 自定义类型名 func(参数列表) 返回值列表
// func timeSpent(inner IntConv) IntConv的意思是什么？
// timeSpent是一个函数，参数是IntConv类型，返回值是IntConv类型
// func timeSpent(inner IntConv) IntConv的格式是什么？
// func 函数名(参数列表) 返回值列表
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

}
