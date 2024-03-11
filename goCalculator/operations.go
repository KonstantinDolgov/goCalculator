package main

import "fmt"

func operations() {
	var a, b int
	var oper string
	fmt.Scan(&a, &oper, &b)
	switch oper {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(sub(a, b))
	case "*":
		fmt.Println(mult(a, b))
	case "/":
		fmt.Println(div(a, b))
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
