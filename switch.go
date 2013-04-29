package main

import "fmt"
import "time"

func main() {
  // basic switch
  i := 2
  fmt.Print("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  // multiple expression using commas
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  // default case
  default:
    fmt.Println("It's a weekday")
  }

  // An alternate way to express if/else
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }
}