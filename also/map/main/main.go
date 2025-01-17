package main

import "fmt"

func main() {
	map_range()
}

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
}
