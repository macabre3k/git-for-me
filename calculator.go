package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Roman struct {
	numeral string
	value   int
}

var RomanNumeral = []Roman{
	{"XC", 90}, {"C", 100}, {"XL", 40}, {"L", 50},
	{"IX", 9}, {"X", 10}, {"IV", 4}, {"V", 5},
	{"I", 1},
}

var RomanValue = []Roman{
	{"C", 100}, {"XC", 90}, {"L", 50}, {"XL", 40},
	{"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4},
	{"I", 1},
}
func RomanNumberToInt(romanNumber string) (result int) {
	count := len(romanNumber)
	runes := []rune(romanNumber)

	i := 0
	for i < count {
		substring := string(runes[i:count])

		var matching Roman

		for _, glyph := range RomanNumeral {
			if strings.HasPrefix(substring, glyph.numeral) {
				matching = glyph
				break
			}
		}
		result += matching.value
		i += len(matching.numeral)
	}
	return result
}
func IntToRomanNumber(number int) (result string) {
	for _, romanGlyph := range RomanValue {
		for number >= romanGlyph.value {
			result += romanGlyph.numeral
			number -= romanGlyph.value
		}
	}
	return result
}
func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

func divi(a int, b int) int {
	return a / b
}

func main() {
	var result int
	var isRoman bool

	fmt.Println("Введите выражение")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	elemets := strings.Fields(text) 

	if len(elemets) > 3 {
		fmt.Println("Ошибка операции")
		return
	}
	if len(elemets) < 3 {
		fmt.Println("Ошибка операции")
		return
	}

	num1, er1 := strconv.Atoi(elemets[0])
	num2, er2 := strconv.Atoi(elemets[2])

	if (er1 == nil && er2 != nil) || (er1 != nil && er2 == nil) {
		fmt.Println("Вывод ошибки, числа либо римские либо арабские")
		return
	}
	if er1 != nil && er2 != nil {
		isRoman = true

		num1 = RomanNumberToInt(elemets[0])
		num2 = RomanNumberToInt(elemets[2])
	}
	if (num1 > 10 || 1 > num1) || (num2 > 10 || 1 > num2) {
		fmt.Println("Числа должны быть от 1 до 10 включительно")
		return
	}
	switch elemets[1] {
	case "+":
		result = sum(num1, num2)
	case "-":
		result = sub(num1, num2)
	case "/":
		result = divi(num1, num2)
	case "*":
		result = multi(num1, num2)
	default:
		println("Ошибка")
		return
	}
if isRoman {
		if result < 1 {
			println("Ошибка")
			return
		}
		println(IntToRomanNumber(result))

	} else {
		println(result)
	}
}
