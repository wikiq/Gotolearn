package main

import (
	"fmt"
	"math"
)

func hybot(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

func fire(a, b int) {
	fmt.Println("fire")
	fmt.Println(a + b)
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	// fmt.Println(hybot(3, 4))

	var f func(a, b int)
	f = fire
	f(10, 20)
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}
