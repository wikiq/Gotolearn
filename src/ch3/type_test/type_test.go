package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr=aPtr+1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

//aPtr := &a的含义是什么？
//1.定义一个变量a，赋值为1
//2.定义一个变量aPtr，赋值为a的地址
//3.打印a的值
//4.打印aPtr的值
//5.打印a的类型
//6.打印aPtr的类型
//7.结束

func TestString(t *testing.T) {
	var s string = "123"
	t.Log("*" + s + "*")
	t.Log(len(s))
	/*	if s == "" {
		t.Log("s is empty")
	}*/
}
