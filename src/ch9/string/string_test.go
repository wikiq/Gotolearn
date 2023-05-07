package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值“”
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	/*s = "\xE4\xBA\xBB\xFF"*/
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //是byte数
	c := []rune(s)
	//为什么rune后len只剩下了1？
	//因为rune是int32，占4个字节，所以len只剩下了1
	t.Log(len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
	//%[1]c %[1]d是什么意思？
	//%[1]c是字符，%[1]d是数字
	//%c是带代表字符，%d是带代表数字，对吗？
	//对
}
