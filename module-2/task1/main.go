package main

import (
	"fmt"
	"strconv"

	"github.com/pbochniak/computator/fibonacci"
)

func main() {
	var inputNumber string
	var number interface{}
	fmt.Printf("Give me a number: ")
	fmt.Scanf("%s", &inputNumber)
	if u, err := strconv.ParseUint(inputNumber, 10, 0); err == nil {
		number = uint(u)
	} else if i, err := strconv.ParseInt(inputNumber, 10, 0); err == nil {
		number = int(i)
	} else if f, err := strconv.ParseFloat(inputNumber, 64); err == nil {
		number = f
	} else {
		fmt.Printf("Error: %q isn't a number.\n", inputNumber)
	}
	if number != nil {
		fmt.Printf("Fibonacci sequence for given number (%v [%T]) equals %v\n", number, number, fibonacci.Fibonacci(number))
	}
}
