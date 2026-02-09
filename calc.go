package main

import "fmt"

func main () {
  var first_no int
  var second_no int
  var operator_int int

  operators := map[int] string{0:"+", 1:"-", 2:"*", 3:"/"}
  fmt.Print("Enter the first number: ")
  fmt.Scanln(&first_no)
  fmt.Print("Enter the second number: ")
  fmt.Scanln(&second_no)
  fmt.Print("Enter operation")
  fmt.Scanln(&operator_int)
  
  op :=operators[operator_int]
  result := calculate(first_no, second_no, op)
  fmt.Println(result)
}


func calculate(a, b int, op string) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        return a / b
    default:
        panic("invalid operator")
    }
}


