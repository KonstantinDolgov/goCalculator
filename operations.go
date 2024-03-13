package main

import (
	"strings"
)

func operations(a, b int) (res int) {
	switch {
	case strings.Contains(text, "+"):
		res = add(a, b)
	case strings.Contains(text, "-"):
		res = sub(a, b)
	case strings.Contains(text, "*"):
		res = mult(a, b)
	case strings.Contains(text, "/"):
		res = div(a, b)
	}
	return
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
