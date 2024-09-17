package main

import (
	"fmt"
)

func main() {
	var number1 int
	var number2 int
	var oper string
	//var Result float32

	fmt.Print("Enter first Number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter second Number: ")
	fmt.Scan(&number2)
	fmt.Print("Choose your operator (+ - * /): ")
	fmt.Scan(&oper)

	switch oper {
	case "+":
		fmt.Println("Result:", number1+number2)
	case "-":
		fmt.Println("Result:", number1-number2)
	case "*":
		fmt.Println("Result:", number1*number2)
	case "/":
		if number2 != 0 {
			fmt.Println("Result:", number1/number2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid operator!")
	}
}
