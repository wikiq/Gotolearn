package main

import "fmt"

func main() {
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

}
