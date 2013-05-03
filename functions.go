package main

import "fmt"

func plus(a int, b int) int {
	return a + b // Go requires explicit returns
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
