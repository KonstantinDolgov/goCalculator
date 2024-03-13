package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var text string

func logic() {
	text, _ = reader.ReadString('\n')
	text = strings.ReplaceAll(strings.TrimSpace(text), " ", "")
	operandsTemp := regexp.MustCompile(`[+\-*/]`).Split(text, -1)
	var operands []string
	for _, v := range operandsTemp {
		if v != "" {
			operands = append(operands, v)
		}
	}
	if len(operands) > 2 {
		panic("Формат математической операции не удовлетворяет заданию " +
			"— два операнда и один оператор (+, -, /, *)")
	}
	if len(operands) < 2 {
		panic("Строка не является математической операцией")
	}

	if (!isRoman(operands[0]) && isRoman(operands[1])) || (isRoman(operands[0]) && !isRoman(operands[1])) {
		panic("Используются одновременно разные системы счисления!")
	}

	if isRoman(operands[0]) && isRoman(operands[1]) {
		res := operations(convertToArabic(operands[0], operands[1]))
		if res < 1 {
			panic("В римской системе нет нуля и отрицательных чисел!")
		}
		fmt.Println(convertToRoman(res))
	}

	if !isRoman(operands[0]) && !isRoman(operands[1]) {
		num1, _ := strconv.Atoi(operands[0])
		num2, _ := strconv.Atoi(operands[1])
		if num1 > 10 || num2 > 10 {
			panic("Операнды больше 10!")
		}
		fmt.Println(operations(num1, num2))
	}
}

func convertToArabic(r1, r2 string) (int, int) {
	if romans[r1] > 10 || romans[r2] > 10 {
		panic("Операнды больше 10!")
	}
	return romans[r1], romans[r2]
}

func convertToRoman(r int) (res string) {
	for key, val := range romans {
		if val == r {
			res = key
		}
	}
	return
}
