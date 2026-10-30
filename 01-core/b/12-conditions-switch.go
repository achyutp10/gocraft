package main

import "fmt"

func main() {
	// switch expression {
	// case condition:

	// }

	var marks int
	fmt.Println("Enter your marks: ")
	fmt.Scanln(&marks)

	switch {
	case marks < 40:
		fmt.Println("Failed")
	case marks >= 40 && marks <= 50:
		fmt.Println("C Graded")
	case marks > 50 && marks <= 65:
		fmt.Println("B Graded")
	case marks > 65:
		fmt.Println("A Graded")
	default:
		fmt.Println("No Marks added")

	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Sunday":
		fmt.Println("Sunday weekend")
	default:
		fmt.Println("Invalid")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than one")
		fallthrough
	case num == 2:
		fmt.Println("Number is two")
	default:
		fmt.Println("Not 2")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("Apple")
	case "banana":
		fmt.Println("Banana")
	default:
		fmt.Println("Unknown")
	}

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
