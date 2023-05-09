package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("A Number:", rand.Intn(10))
	//rand.Intn(10) 生成一个0-10的随机数
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
