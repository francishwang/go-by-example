package main

import "fmt"
import "math"

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n
  fmt.Println(d)

  // Explicit cast: A numeric constant has no type
  fmt.Println(int64(d))

  // implicit cast: math.Sin expects a float64
  fmt.Println(math.Sin(n))
}
