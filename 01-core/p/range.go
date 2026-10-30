package main

import "fmt"

func main() {

	// var mapVariable map[keyType] valueType
	// mapVariable = make(map[keyType] valueType)

	// usng a map literal
	// mapVariable = map[keyType] valueType {
	// 	key1: value1,
	// 	key2: value2
	// }

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)

	// a := [5]int{1, 2, 3, 4, 5}

	// array
	// var numbers [5]int
	// fmt.Println(numbers)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	map1 := map[string]interface{}{
		"ram": map[string]string{
			"job":  "Marketing",
			"city": "Brt",
		},
		"shyam": 12,
		"sita":  12,
		"details": map[string]string{
			"job":  "Marketing",
			"city": "Brt",
		},
	}

	k, v := map1["ram"].(map[string]string)
	if v {
		fmt.Println(k["job"])
	}

	fmt.Println(map1["ram"])

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	slice1 := []int{9, 8, 7}
	slice2 := [][]int{
		{9, 8, 7},
		{1, 2, 3},
		{6, 7, 9},
	}

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	for i, v := range slice2 {
		fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}

	var noArray [5]int = [5]int{1, 2, 3, 4, 5}
	var noArray2D [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, v := range noArray {
		fmt.Println(i, v)
	}

	for _, v := range noArray2D {
		// fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}

}
