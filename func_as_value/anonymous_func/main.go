package main

import "fmt"

func assert(fn func(int, int) int, a int, b int, expected int) bool {
	return fn(a, b) == expected
}

func main() {
	fmt.Println(assert(
		func(a int, b int) int {
			return a + b
		},
		1, 2, 3,
	))
}
