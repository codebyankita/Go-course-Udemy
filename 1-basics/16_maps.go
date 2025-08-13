package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	// fmt.Println(value)
	fmt.Println("Is a value associated with key1:", ok)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	fmt.Println("myMap length is", len(myMap))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}