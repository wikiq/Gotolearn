package fib

import (
	"fmt"
	"testing"
)

var a int

//上面这串代码的含义是什么？
//1.定义两个变量a和b，分别赋值为1
//2.打印a的值
//3.循环5次
//4.打印b的值
//5.将a的值赋值给tmp
//6.将b的值赋值给a
//7.将tmp+a的值赋值给b
//8.结束循环

func TestFibList(t *testing.T) {
	/*	var a int = 1
		var b int = 1*/
	/*	var (
		a int = 1
		b int = 1
	)*/
	a = 1
	b := 1
	fmt.Println(a, "")
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Println("", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	/*	tmp := a
		a = b
		b = tmp*/
	a, b = b, a
	t.Log(a, b)
}
