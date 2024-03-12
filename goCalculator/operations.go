package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func operations() {
	var a, b int
	var oper string
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//text = strings.TrimSpace(text)
	//toNumber, _ := strconv.Atoi(text)
	operands := regexp.MustCompile(`[+\-*/]`).Split(text, 0)
	if len(operands) > 2 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
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
