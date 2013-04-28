package main

import "fmt"

func main() {
  // basic branching with if & else:
  // don't need parentheses around conditions
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  // without else
  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  // if/else-if/else with var declaration
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }

  // Note: Go has no ternary if (e.g. ?:)
}
