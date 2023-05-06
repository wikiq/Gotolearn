package slice_test

import (
	"testing"
)

func TestSliceinit(t *testing.T) {
	/*var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	//append的含义是追加，如果追加的元素超过了cap，那么会重新分配更大的底层数组
	//为什么追加之后的len和cap都变成了1？
	//因为append的时候，如果超过了cap，那么会重新分配更大的底层数组，然后把原来的元素复制过去
	//请写出详细过程append中的底层数组是如何变化的
	s0 = append(s0, 2)
	t.Log(len(s0), cap(s0))
	//这个len和cap的是什么
	//len是指slice中的元素个数，cap是指slice中的底层数组的容量
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))*/

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
	s2 = append(s2, 2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	/*	t.Log(s2[5])*/
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
		t.Log(i)
		t.Log(s[0:(len(s))])
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	year := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	Q2 := year[3:6]
	//为什么Q2的cap是9？
	//因为Q2是year[3:6]，所以Q2的cap是year[3:]的cap，而year[3:]的cap是9
	//具体解释一下为什么year[3:]的cap是9？
	//因为year[3:]是从year[3]开始的，而year[3]是four，所以year[3:]的cap是从four开始的，也就是nine，所以year[3:]的cap是9
	//所以cap=9就代表从four到twelve，一共9个元素，这么理解对吗？
	//对
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	//所以此处的cap=7是因为six到twelve一共7个元素，对吗？
	//对
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(len(Q2), cap(Q2))
	//six变成了Unknown，是因为Q2和summer共享了底层数组，所以summer[0]的修改会影响到Q2
	//因为原本summer[0]=“six”，现在summer[0]=“Unknown”，所以Q2[0]=“Unknown”
	t.Log(year)
}

//数组是如何比较的？
//数组是值类型，所以数组是通过值来比较的
//切片是如何比较的？
//切片是引用类型，所以切片是通过引用来比较的
//切片的比较是不安全的，所以不要用切片来比较
//切片只能和nil比较

func TestSliceCopare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//t.Log(a == b)
	//切片是引用类型，所以切片是通过引用来比较的，所以切片的比较是不安全的，所以不要用切片来比较
	//切片只能和nil比较
	t.Log(a == nil)
	t.Log(b == nil)
	/*	if a == b {
		t.Log("equal")
	}*/
}
