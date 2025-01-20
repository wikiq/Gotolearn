package main

import "fmt"

func break_test() {
OuterLoop:
	//这个OuterLoop是标签,可以在break的时候指定跳出到哪个标签break_test
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			//第一次执行时，i=0,j=2时，跳出外层循环
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}

func continue_test() {
OuterLoop2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop2
			}
		}
	}
}

func main() {
	break_test()
	continue_test()

}
