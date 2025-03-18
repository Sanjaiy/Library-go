package main

import (
	"flag"
	"fmt"
)

// "bufio"
// "fmt"
// "os"
// "strconv"

// Input based calculator //
// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("What would you like to do? (add, subract, multiply, divide)")
// 	input, _ := reader.ReadString('\n')

// 	fmt.Println("Enter First Number:")
// 	firstNumber, _ := reader.ReadString('\n')

// 	num1, err := strconv.ParseFloat(firstNumber[:len(firstNumber)-1], 64)
// 	if err != nil {
// 		fmt.Println("Please enter a valid number")
// 	}

// 	fmt.Println("Enter Last Number:")
// 	lastNumber, _ := reader.ReadString('\n')

// 	num2, err := strconv.ParseFloat(lastNumber[:len(lastNumber)-1], 64)
// 	if err != nil {
// 		fmt.Println("Please enter a valid number")
// 	}

// 	switch input[:len(input)-1] {
// 	case "add":
// 		fmt.Println("Result: ", num1 + num2)
// 	case "subtract":
// 		fmt.Println("Result: ", num1 - num2)
// 	case "multiply":
// 		fmt.Println("Result: ", num1 * num2)
// 	case "divide":
// 		fmt.Println("Result: ", num1 / num2)
// 	default:
// 		fmt.Println("Invalid Input")
// 	}
// }

// Flag based calculator //
func main() {
	operation := flag.String("operation", "add", "operation to perform")
	num1 := flag.Float64("num1", 0, "first number")
	num2 := flag.Float64("num2", 0, "second number")

	flag.Parse()

	switch *operation {
	case "add":
		fmt.Println("Result: ", *num1 + *num2)
	case "subtract":
		fmt.Println("Result: ", *num1 - *num2)
	case "multiply":
		fmt.Println("Result: ", *num1 * *num2)
	case "divide":
		if *num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed")
			return
		}
		fmt.Println("Result: ", *num1 / *num2)
	default:
		fmt.Println("Invalid Input")
	}
}