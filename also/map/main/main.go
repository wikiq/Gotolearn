package main

import (
	"fmt"
	"sort"
	"sync"
)

// func main() {
// 	map_test()
// 	map_range_two()
// 	scence_map()
// 	m := make(map[int]int)
// 	go func() {
// 		for {
// 			m[1] = 1
// 		}
// 	}()
// 	go func() {
// 		for {
// 			_ = m[1]
// 		}
// 	}()
// 	for {
// 	}
// }

func map_test() {
	var mapList map[string]int
	var mapCreated map[string]float64
	var maspAssigned map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	maspAssigned = mapList
	mapCreated = make(map[string]float64)
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	maspAssigned["two"] = 3
	fmt.Printf("Map Literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map Created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map Assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map Literal at \"ten\" is: %d\n", mapList["ten"])

	// noteFrequency := map[string]float32{
	// 	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	// 	"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	// fmt.Println(noteFrequency)
}

func map_range() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for key, value := range scene {
		fmt.Println("Key:", key, "Value:", value)
	}
	for key := range scene {
		fmt.Println("Key:", key)
	}
}

func map_range_two() {
	scence := make(map[string]int)

	scence["route"] = 66
	scence["brazil"] = 4
	scence["china"] = 960

	var scence2 []string
	for key := range scence {
		scence2 = append(scence2, key)
	}
	sort.Strings(scence2)
	fmt.Println(scence2)
	delete(scence, "route")
	fmt.Println(scence)
}

func scence_map() {
	var scence3 sync.Map
	scence3.Store("greece", 97)
	scence3.Store("london", 100)
	scence3.Store("egypt", 200)

	fmt.Println(scence3.Load("london"))
	scence3.Delete("london")
	// fmt.Println(scence3.Load("london"))
	scence3.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

func main() {
	scence_map()
}
