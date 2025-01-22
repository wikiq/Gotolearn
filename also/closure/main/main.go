package main

import (
	"fmt"
)

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {

	// 返回一个闭包
	return func() int {

		// 累加
		value++

		// 返回一个累加值
		return value
	}
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func finbionacci(n int) (res int) {
	if n < 2 {
		res = n
	} else {
		res = finbionacci(n-1) + finbionacci(n-2)
	}
	return
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	// myfunc(1, 2, 3, 4, 5)
	// myfunc(2, 3, 4)
	// var v1 int = 1
	// var v2 int64 = 234
	// var v3 string = "hello"
	// var v4 float32 = 1.234
	// MyPrintf(v1, v2, v3, v4)

	// fmt.Println("defer begin")
	// // 将defer放入延迟调用栈
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// // 最后一个放入, 位于栈顶, 最先调用
	// defer fmt.Println(3)
	// fmt.Println("defer end")
	result := 0
	for i := 0; i < 10; i++ {
		result = finbionacci(i)
		fmt.Printf("Fibonacci(%d) is:%d\n", i, result)
	}
	var i int = 15
	fmt.Println("i=", i, "Factorial is ", Factorial(uint64(i)))
}
