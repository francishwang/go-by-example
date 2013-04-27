package main

import "fmt"

func main() {
  var a string = "initial"
  fmt.Println(a)

  // multiple vars
  var b, c int = 1, 2
  fmt.Println(b, c)

  // type inference
  var d = true
  fmt.Println(d)

  // default value: 0
  var e int
  fmt.Println(e)

  // shorthand syntax for: var f string = "short"
  f := "short"
  fmt.Println(f)
}
