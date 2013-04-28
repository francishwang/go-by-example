package main

import "fmt"

func main() {
  // single condition
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // initial/condition/after
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // just loop
  for {
    fmt.Println("loop")
    break
  }
}
