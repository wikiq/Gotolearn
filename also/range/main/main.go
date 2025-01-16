package main

import "fmt"

func main() {
	// slice := []int{10, 20, 30, 40}
	// // 迭代每一个元素，并显示其值
	// for index, value := range slice {
	// 	fmt.Printf("Index: %d Value: %d\n", index, value)
	// }

	// gopath := os.Getenv("GOPATH")
    // fmt.Println("GOPATH:", gopath)

	// slice := []int{10, 20, 30, 40}
	// // 迭代每个元素，并显示值和地址
	// for index, value := range slice {
	// 	fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	// 	//给我详细解释一下上面这个语句

	// }
	// testSlice := []int{10, 20, 30, 40, 50}
	// for _, value := range testSlice {
	// 	fmt.Printf("Value: %d\n", value)
	// }
	slice := []int{10, 20, 30, 40, 50,60, 70, 80, 90, 100}
	// 从第三个元素开始迭代每个元素
	for index := 3; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}
