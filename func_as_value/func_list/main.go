package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a int, b int) (result int, operation string) {
	result = a + b
	operation = "a + b"
	return
}

func sub(a int, b int) (result int, operation string) {
	result = a - b
	operation = "a - b"
	return
}

func mul(a int, b int) (result int, operation string) {
	result = a * b
	operation = "a * b"
	return
}

func div(a int, b int) (result int, operation string) {
	result = a / b
	operation = "a / b"
	return
}

func getArgs(args []string) (a int, b int) {
	if len(os.Args) != 3 {
		fmt.Println("Please specify two args.")
		os.Exit(1)
	}

	a, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("cannot read args as int: %s\n", err)
		os.Exit(1)
	}

	b, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("cannot read args as int: %s\n", err)
		os.Exit(1)
	}

	return
}

func main() {
	a, b := getArgs(os.Args)

	var operations = [4]func(int, int) (int, string){
		add, sub, mul, div,
	}

	fmt.Printf("calculate for (a, b) = (%d, %d)\n", a, b)

	for i := 0; i < len(operations); i++ {
		result, operation := operations[i](a, b)
		fmt.Printf("%s = %d\n", operation, result)
	}
}
