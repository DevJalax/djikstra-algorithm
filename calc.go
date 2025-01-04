package main

import (
	"fmt"
)

func calc(c int, d int, op string) int {
	switch op {
	case "+":
		return c + d
	case "-":
		return c - d
	case "*":
		return c * d
	case "/":
		return c / d
	default:
		fmt.Println("wrong choice")
		return 0
	}
}

func main() {
	fmt.Println("Enter two numbers and operation operator")
	var (
		a      int
		b      int
		result int
		str    string
	)
	fmt.Scanln(&a, &b, &str) // input
	result = calc(a, b, str)
	fmt.Println(result)
}
