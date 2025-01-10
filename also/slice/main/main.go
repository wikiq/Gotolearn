package main

import (
	"fmt"
)

func main() {
	// shuzu2()
	slice3()
}

func shuzu()  {
	a := [3]int{5,10,15}
	// fmt.Println(a)
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }
	fmt.Println(a[len(a)-1])
}
func shuzu2()  {
	// var array [4][2][3]int
	// fmt.Println(array)
	// var array [4][2]int
	// array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// fmt.Println(array)
	array1 := [2][2]int{{1, 2}, {3, 4}}
	var array3 [2]int = array1[1]
	fmt.Println(array3)
	vlaue := array1[1][1]
	fmt.Println(vlaue)
}
func slice1()  {
	var strList []string
	var numList []int
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)
}
func slice2()  {
	var at []int
	at = append(at, 1, 2, 3)
	fmt.Println(at)
	fmt.Println(cap(at))
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}

func slice3()  {
	var a = []int{1,2,3}
	// fmt.Println(a)
	// a = append([]int{0}, a...) // 在开头添加1个元素
	// fmt.Println(a)
	a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
	fmt.Println(a)

	a = append(a[:3],append([]int{5},a[1:]...)...)
	fmt.Println(a)
	a = append(a[:1],append([]int{1,2,3},a[3:]...)...)
	fmt.Println(a)
}
