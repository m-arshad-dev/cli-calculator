package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to CLI Calculator!")
	fmt.Println("You can use +, -, *, /, %, ^ operators or type 'exit' to quit.")

	var num1, num2 float64
	var operator string

	for {
		fmt.Print("\nEnter first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter operator (+, -, *, /, %, ^) or 'exit' to quit: ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("Exiting the calculator. Goodbye!")
			return
		}

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("Result: %.2f\n", num1/num2)
			} else {
				fmt.Println("Error: Division by zero")
			}
		case "%":
			if num2 != 0 {
				fmt.Printf("Result: %.2f\n", float64(int(num1)%int(num2)))
			} else {
				fmt.Println("Error: Division by zero")
			}
		case "^":
			fmt.Printf("Result: %.2f\n", math.Pow(num1, num2))
		default:
			fmt.Println("Error: Unknown operator")
		}
	}
}
