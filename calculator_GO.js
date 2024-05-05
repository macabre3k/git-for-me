package main

import (
    "fmt"
    "strings" // добавлен импорт пакета strings
)

// Карта для преобразования римских чисел в арабские
var romanToArabic = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Карта для преобразования арабских чисел в римские
var arabicToRoman = map[int]string{
    1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
    20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

// Функция для преобразования римских чисел в арабские
func romanToArabicConverter(roman string) (int, error) {
    result := 0
    for i := 0; i < len(roman); i++ {
        current := romanToArabic[string(roman[i])]
        next := 0
        if i+1 < len(roman) {
            next = romanToArabic[string(roman[i+1])]
        }
        if next > current {
            result += next - current
            i++
        } else {
            result += current
        }
    }
    return result, nil
}

// Функция для преобразования арабских чисел в римские
func arabicToRomanConverter(arabic int) string {
    var roman strings.Builder
    for arabic >= 100 {
        roman.WriteString("C")
        arabic -= 100
    }
    if arabic >= 90 {
        roman.WriteString("XC")
        arabic -= 90
    }
    if arabic >= 50 {
        roman.WriteString("L")
        arabic -= 50
    }
    if arabic >= 40 {
        roman.WriteString("XL")
        arabic -= 40
    }
    for arabic >= 10 {
        roman.WriteString("X")
        arabic -= 10
    }
    if arabic >= 9 {
        roman.WriteString("IX")
        arabic -= 9
    }
    if arabic >= 5 {
        roman.WriteString("V")
        arabic -= 5
    }
    if arabic >= 4 {
        roman.WriteString("IV")
        arabic -= 4
    }
    for arabic >= 1 {
        roman.WriteString("I")
        arabic -= 1
    }
    return roman.String()
}

func main() {
    var num1, num2 float64
    var operation string

    fmt.Print("Введите первое число: ")
    fmt.Scanln(&num1)

    fmt.Print("Введите операцию (+, -, *, /): ")
    fmt.Scanln(&operation)

    fmt.Print("Введите второе число: ")
    fmt.Scanln(&num2)

    var result float64

    switch operation {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 != 0 {
            result = num1 / num2
        } else {
            fmt.Println("Ошибка: деление на ноль невозможно")
            return
        }
    default:
        fmt.Println("Ошибка: неверная операция")
        return
    }

    fmt.Printf("Результат: %.2f %s %.2f = %.2f", num1, operation, num2, result)
}
