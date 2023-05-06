package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Tursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
	pending
)

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

//a&Readable == Readable的意思是什么？
//a&Readable == Readable的意思是a和Readable进行按位与运算，如果结果等于Readable，那么就是true，否则就是false
//a&Readable如何计算？
//a&Readable的计算过程如下：
//1.将a转换为二进制数：0001
//2.将Readable转换为二进制数：0001
//3.将两个二进制数进行按位与运算：0001
//4.将运算结果转换为十进制数：1
//5.将1和Readable进行比较，如果相等，那么就是true，否则就是false
//0001&0111=0001
//0001&0010=0000

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Tursday, Friday, Saturday, Sunday)
}
