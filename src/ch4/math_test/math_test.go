package math_test

import "testing"

func TestCompareArray(t *testing.T) {

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)

}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(a&0001, a&0010, a&0100)
}

//a&Readable == Readable, a&Writable == Writable, a&Executable == Executable的计算过程是什么
//a&Readable == Readable, a&Writable == Writable, a&Executable == Executable的计算过程如下：
//1.将a转换为二进制数：0111
//2.将Readable转换为二进制数：0001
//3.将Writable转换为二进制数：0010
//4.将Executable转换为二进制数：0100
//5.将两个二进制数进行按位与运算：0111&0001=0001，0111&0010=0010，0111&0100=0100
//6.将运算结果转换为十进制数：1，2，4
//7.将1，2，4和Readable，Writable，Executable进行比较，如果相等，那么就是true，否则就是false
//8.结束

//&^ 如果右边的位置为1，无论左边是0或是1，结果都是0
//&^ 如果右边的位置为0，结果都是左边的值

//a = a &^ Readable
//a = a &^ Writable的计算过程
//1.将a转换为二进制数：0111
//2.将Readable转换为二进制数：0001
//3.将Writable转换为二进制数：0010
//4.将两个二进制数进行按位与运算：0111&^0001=0110，0110&^0010=0100
//5.将运算结果转换为十进制数：4
//6.将4和Readable，Writable进行比较，如果相等，那么就是true，否则就是false
//7.结束
