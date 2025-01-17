package main

import "fmt"

func switch_test() {
	switch {
	case 1 == 1:
		println("1")
	case 1 == 2:
		println("2")
	default:
		println("3")
	}
}

func main() {
	// switch_test()
	gototest2()
}

func gototest() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("done")
}

func gototest2() {
	// err := firstCheckError()
	// if err != nil {
	// 	fmt.Println(err)
	// 	exitProcess()
	// 	return
	// }
	// err = secondCheckError()

	// if err != nil {
	// 	fmt.Println(err)
	// 	exitProcess()
	// 	return
	// }
	// fmt.Println("done")
	// 	err := firstCheckError()
	// 	if err != nil {
	// 		goto onExit
	// 	}

	// 	err = secondCheckError()

	// 	if err != nil {
	// 		goto onExit
	// 	}

	// 	fmt.Println("done")

	// 	return

	// onExit:
	//
	//	fmt.Println(err)
	//	exitProcess()
}
