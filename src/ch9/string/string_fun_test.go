package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "_"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	//strconv是个库吗？
	//是
	//Itoa是个函数吗？
	//是
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
	//Itoa是什么意思？
	//int to string
	//Atoi是什么意思？
	//string to int
	//为什么要有这两个函数？
	//因为有时候我们需要把int转换成string，有时候我们需要把string转换成int
}
