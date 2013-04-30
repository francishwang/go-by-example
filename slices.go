package main

import "fmt"

func main() {
  // empty slice with 3 length (initially zero-valued)
  s := make([]string, 3)
  fmt.Println("emp:", s)

  // set & get just like with arrays
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])
  fmt.Println("len:", len(s))

  // append
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  // copy c from s in same length
  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  // gets a slice of the elements s[2], s[3] & s[4]
  // cf. s[5] is excluded
  l := s[2:5]
  fmt.Println("sl1:", l)

  // up to s[5] (also excluding s[5])
  l = s[:5]
  fmt.Println("sl2:", l)

  // up from s[2] (including s[2])
  l = s[2:]
  fmt.Println("sl3:", l)

  // declare & initialize a slice at once
  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  // multi-dimensional slice
  // The length of the inner slices can vary, unlike ararys
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
