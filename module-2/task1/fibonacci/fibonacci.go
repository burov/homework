package fibonacci

import "fmt"

func Fibonacci(n interface{}) interface{} {
	switch typed_n := n.(type) {
	case uint:
		return fibonnaciUint(typed_n)
	case int:
		return fibonnaciInt(typed_n)
	case float64:
		return fibonnaciFloat(typed_n)
	default:
		return fmt.Errorf("unknown fibonacci input argument type: %T", typed_n)
	}
}

func fibonnaciUint(n uint) uint {
	var previous uint = 0
	var current uint = 1
	if n < 1 {
		return previous
	} else if n == 1 {
		return current
	}
	for i := uint(2); i <= n; i++ {
		current, previous = current+previous, current
	}
	return current
}

func fibonnaciInt(n int) int {
	var previous int = 0
	var current int = 1
	if n < 1 {
		return previous
	} else if n == 1 {
		return current
	}
	for i := int(2); i <= n; i++ {
		current, previous = current+previous, current
	}
	return current
}

func fibonnaciFloat(n float64) float64 {
	var previous float64 = 0
	var current float64 = 1
	if n < 1 {
		return previous
	} else if n == 1 {
		return current
	}
	for i := float64(2); i <= n; i++ {
		current, previous = current+previous, current
	}
	return current
}
