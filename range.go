package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// iterate array over all elements
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// iterate with index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// iterate map over key/values pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate string over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
