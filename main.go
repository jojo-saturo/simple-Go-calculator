package main

import (
	"fmt"
)

func main() {
	var operation string
	var num1, num2 float64

	fmt.Print("Enter operation (+, -, *, /) ")
	fmt.Scanln(&operation)

	fmt.Print("Enter first number ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number ")
	fmt.Scanln(&num2)

	switch operation {
	case "+":
		fmt.Printf("Result: %g\n", num1+num2)
	case "-":
		fmt.Printf("Result: %g\n", num1-num2)
	case "*":
		fmt.Printf("Result: %g\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %g\n", num1/num2)
		} else {
			fmt.Println("Error: Undefined")
		}
	default:
		fmt.Println("Error: Invalid operation")
	}
}
