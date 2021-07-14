package main

import (
	"fmt"
	"honestit/computator/fibonacci"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Provide a number for fibonnaci function: ")
	fmt.Scanf("%s", &input)

	n, err := cast(input)
	if err == nil {
		fib := fibonacci.Fibonacci(n)
		fmt.Printf("Fibonacci for %v = %v", n, fib)
	} else {
		fmt.Printf("Something wrong with your input: %q", input)
	}

}

func cast(input string) (interface{}, error) {
	if v, err := strconv.ParseInt(input, 0, 0); err == nil {
		return int(v), nil
	} else if v, err := strconv.ParseUint(input, 0, 0); err == nil {
		return uint(v), nil
	} else if v, err := strconv.ParseFloat(input, 64); err == nil {
		return float64(v), nil
	} else {
		return nil, fmt.Errorf("not a number")
	}
}
