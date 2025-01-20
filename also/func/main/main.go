package main

import (
	"flag"
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

func visit2(ulist []int, f func(int)) {
	for _, v := range ulist {
		f(v)
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	// fmt.Println(hybot(3, 4))
	// var f func(a, b int)
	// f = fire
	// f(10, 20)
	// visit([]int{1, 2, 3, 4}, func(v int) {
	// 	fmt.Println(v)
	// })
	// visit2([]int{1, 2, 3, 4, 5, 6}, func(v int) {
	// 	fmt.Println(v)
	// })

	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"runner": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f() //该函数来源于map，所以需要用括号调用，上文f()在map中是值
	} else {
		fmt.Println("skill not found")
	}
}
