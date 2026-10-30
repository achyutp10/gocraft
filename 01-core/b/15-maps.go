package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType] valueType
	// mapVariable = make(map[keyType] valueType)

	// usng a map literal
	// mapVariable = map[keyType] valueType {
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 19
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 91
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)
	// clear(myMap)
	// fmt.Println(myMap)

	// value, unknownValue := myMap["key1"]
	// fmt.Println(value)
	// fmt.Println(unknownValue)

	_, ok := myMap["key1"]
	// fmt.Println(value)
	if ok {
		fmt.Println("Is a value associated with key1: ", ok)
	} else {
		fmt.Println("No value associated with key1: ")
	}

	// _, unknownValue := myMap["key1"]
	// // fmt.Println(value)
	// fmt.Println("Is a value associated with key1: ", unknownValue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap, myMap2) {
		fmt.Println("both map are equal")
	}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("both map are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is iniialized to nil value")
	} else {
		fmt.Println("The map is not iniialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	fmt.Println(len(myMap4))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
