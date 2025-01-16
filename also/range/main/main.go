package main

import "fmt"
import "os"

func main() {
	// slice := []int{10, 20, 30, 40}
	// // 迭代每一个元素，并显示其值
	// for index, value := range slice {
	// 	fmt.Printf("Index: %d Value: %d\n", index, value)
	// }

	gopath := os.Getenv("GOPATH")
    fmt.Println("GOPATH:", gopath)

	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
}
}
