package random

import "fmt"

var (
	num1     float32
	num2     float32
	operator string
)

func Calculator() {
	takeInput()
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid!")
		Calculator()
	}
}

func takeInput() {
	fmt.Print("First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Operation: ")
	fmt.Scan(&operator)
	fmt.Print("Second Number: ")
	fmt.Scan(&num2)
}
