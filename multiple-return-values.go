package main

import "fmt"

func vals() (int, int) {
  return 3, 7 // return multiple values
}

func main() {
  // multiple assignment
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  // get a subset with blank identifier _
  _, c := vals()
  fmt.Println(c)
}
