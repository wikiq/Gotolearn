package main

import "fmt"

func main() {
	// seq := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
    // // 指定删除位置
    // index := 3
    // // 查看删除位置之前的元素和之后的元素
    // fmt.Println(seq[:index], seq[index+1:])
    // // 将删除点前后的元素连接起来
    // seq = append(seq[:index], seq[index+3:]...)
    // fmt.Println(seq)

    // 声明一个二维整型切片并赋值
    slice := [][]int{{10,20,50}, {100, 200}}
    // 为第一个切片追加值为 20 的元素
    slice[1] = append(slice[1], 40)
    fmt.Println(slice)
}
