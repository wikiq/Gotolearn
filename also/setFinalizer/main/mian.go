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
	for i := 0; i < 3001; i++ {
		fmt.Println(i, ".", "i like you")

	}
}
