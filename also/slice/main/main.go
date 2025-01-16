package main

import (
	"fmt"
)

func main() {
	// shuzu2()
	// slice3()

	 // 设置元素数量为1000
	 const elementCount = 1000
	 // 预分配足够多的元素切片
	 srcData := make([]int, elementCount)
	 // 将切片赋值
	 for i := 0; i < elementCount; i++ {
		 srcData[i] = i
	 }
	 // 引用切片数据
	 refData := srcData
	 // 预分配足够多的元素切片
	 copyData := make([]int, elementCount)
	 // 将数据复制到新的切片空间中
	 copy(copyData, srcData)
	 // 修改原始数据的第一个元素
	 srcData[0] = 999
	 // 打印引用切片的第一个元素
	 fmt.Println(refData[0])
	 // 打印复制切片的第一个和最后一个元素
	 fmt.Println(copyData[0], copyData[elementCount-1])
	 // 复制原始数据从4到6(不包含)
	 copy(copyData, srcData[4:6])
	 for i := 0; i < 5; i++ {
		 fmt.Printf("%d ", copyData[i])
	 }
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
	a = append(a[:1],append([]int{1,2,3},a[1:]...)...)
	fmt.Println(a)
}
