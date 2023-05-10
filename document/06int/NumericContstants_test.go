package _6int

import (
	"fmt"
	"testing"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
	//<<是什么意思？
	//<<是左移运算符，表示将1的二进制表示向左移动100位，即乘以2的100次方。
	//>>是右移运算符，表示将1的二进制表示向右移动99位，即除以2的99次方。
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func TestNumberic(t *testing.T) {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
