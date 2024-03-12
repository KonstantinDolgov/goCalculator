package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func operations() {
	var a, b int
	//var oper string
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//text = strings.TrimSpace(text)
	text = strings.ReplaceAll(strings.TrimSpace(text), " ", "")
	operandsTemp := regexp.MustCompile(`[+\-*/]`).Split(text, -1)
	var operands []string
	for _, v := range operandsTemp {
		if v != "" {
			operands = append(operands, v)
		}
	}
	fmt.Println(operands)
	fmt.Println(len(operands))
	if len(operands) > 2 {
		panic("Формат математической операции не удовлетворяет заданию " +
			"— два операнда и один оператор (+, -, /, *)")
	}
	if len(operands) < 2 {
		panic("Строка не является математической операцией")
	}
	//oper = "+"
	//if strings.Contains(text, oper) {
	//	fmt.Println("plus")
	//}
	switch {
	case strings.Contains(text, "+"):
		fmt.Println(add(a, b))
	case strings.Contains(text, "-"):
		fmt.Println(sub(a, b))
	case strings.Contains(text, "*"):
		fmt.Println(mult(a, b))
	case strings.Contains(text, "/"):
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
