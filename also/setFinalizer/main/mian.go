// package main

// import (
// 	"log"
// 	"runtime"
// 	"time"
// )

// type Road int

// func findRoad(r *Road) {

// 	log.Println("road:", *r)
// }

// func entry() {
// 	var rd Road = Road(999)
// 	r := &rd

// 	runtime.SetFinalizer(r, findRoad)
// }

// func main() {

// 	entry()

// 	for i := 0; i < 10; i++ {
// 		time.Sleep(time.Second)
// 		runtime.GC()
// 	}

// }
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i+1, ".", "i like you")
	}
	printMultiplicationTable()
}

// package main

// import (
//     "fmt"
// )

// func printMultiplicationTable() {
//     for i := 1; i <= 9; i++ {
//         for j := 1; j <= i; j++ {
//             fmt.Printf("%d * %d = %d\t", j, i, i*j)
//         }
//         fmt.Println()
//     }
// }

//	func main() {
//	    printMultiplicationTable()
//	}
func printMultiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
