package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {

  str1, operator, str2 := input()

  detectstr1, err1 := isSting(str1)

  if err1 != nil {
    panic(err1)
  }

  detectstr2, err2 := isSting(str2)

  if err2 != nil {
    panic(err2)
  }

  switch operator {
  case "+":
    sum(detectstr1, detectstr2, str1, str2)
  case "-":
    sub(detectstr1, detectstr2, str1, str2)
  case "*":
    mult(detectstr1, detectstr2, str1, str2)
  case "/":
    divide(detectstr1, detectstr2, str1, str2)
  default:
    panic("Не найден оператор (+, -, *, /)")
  }
}

func divide(d1, d2 bool, s1, s2 string) {

  if !(d1 && !d2) {
    panic("Ошибка mult")
  }

  s1 = s1[1 : len(s1)-1]
  sizeS1 := len(s1)
  sizeS2, _ := strconv.Atoi(s2)

  if sizeS2 == 0 {
    panic("Деление на ноль")
  }

  resultSize := sizeS1 / sizeS2

  printResult(s1[:resultSize])
}

func printResult(result string) {
  if len(result) > 40 {
    fmt.Print("\"", result[:40], "...\"")
  } else {
    fmt.Print("\"", result, "\"")
  }
}

func sub(d1, d2 bool, s1, s2 string) {

  if !d1 || !d2 {
    panic("Ошибка sum")
  }

  s1 = s1[1 : len(s1)-1] // получение строки без ковычек
  s2 = s2[1 : len(s2)-1]

  startIndex := strings.Index(s1, s2)

  if startIndex == -1 {
    fmt.Print(s1)
  }

  endIndex := startIndex + len(s2)

  result := s1[:startIndex] + s1[endIndex:]

  printResult(result)
}

func sum(d1, d2 bool, s1, s2 string) {

  if !d1 || !d2 {
    panic("Ошибка sum")
  }

  s1 = s1[1 : len(s1)-1] // получение строки без ковычек
  s2 = s2[1 : len(s2)-1]

  printResult(s1 + s2)
}

func mult(d1, d2 bool, s1, s2 string) {
  if !(d1 && !d2) {
    panic("Ошибка mult")
  }

  s1 = s1[1 : len(s1)-1]

  var result string

  num, _ := strconv.Atoi(s2)
  for i := 0; i < num; i++ {
    result += s1
  }

  printResult(result)
}

func input() (string, string, string) {

  reader := bufio.NewReader(os.Stdin)
  s, _ := reader.ReadString('\n')

  result := make([]string, 3)
  temp := ""

  flag := false

  counter := 0
  for i := 0; i < len(s); i++ {

    if s[i] == '"' {
      flag = !flag
    }

    if s[i] == ' ' || s[i] == '\n' {
      if !flag || s[i] == '\n' {
        result[counter] = temp
        temp = ""
        counter++
        continue
      }
    }

    temp += string(s[i])
  }

  return result[0], result[1], result[2]
}

func isSting(a string) (bool, error) {

  if a[0] == a[len(a)-1] && a[0] == '"' {

    if len(a)-2 >= 10 {
      return false, errors.New("Большое количество символов в строке!")
    }
    return true, nil
  }

  num, err := strconv.Atoi(a)
  if err != nil {
    return false, errors.New("Ошибка")
  }

  if num < 1 || num > 10 {
    return false, errors.New("Число должно быть в пределах от 1 до 10")
  }

  return false, nil
}

  
  

 
   
