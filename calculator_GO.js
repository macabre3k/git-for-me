package main
 
import (
    "fmt"
    "strconv"
    "strings" 
)
 // Карта для преобразования римских чисел в арабские
var romanToArabic = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}
 // Карта для преобразования арабских чисел в римские
var arabicToRoman = map[int]string{
    1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}
 // Функция для преобразования римских чисел в арабские
func romanToArabicConverter(roman string) int {
    result := 0
    for i := 0;  i < len(roman); i++ { 
        current  := romanToArabic[string(roman[i])] //текущий = current
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
    return result
}
 
// Функция для преобразования арабских чисел в римские
func arabicToRomanConverter(arabic int) string {
    var roman strings.Builder
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
 func examination(num string) float64 {
    _, ok := romanToArabic[num]
    if ok {
        return float64(romanToArabicConverter(num))
    } else {
        i, err := strconv.Atoi(num)
        if err != nil {
            panic(err)
        }
        return float64(i)
 
    }
}
 
func main() {
    var inputNum1, inputNum2 string
    var result float64
    var operation string
 
    fmt.Print("Введите пример")
    fmt.Scanf("%s %c %s", &inputNum1, &operation, &inputNum2)
    num1 := examination(inputNum1)
    num2 := examination(inputNum2)
    
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
    fmt.Printf("Результат: %s %c %s = %.2f\n", inputNum1, operation, inputNum2, result)
}
