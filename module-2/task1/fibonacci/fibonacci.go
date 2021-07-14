package fibonacci

func Fibonacci(n interface{}) interface{} {
	switch v := n.(type) {
	case uint:
		return fibonnaciUint(v)
	case int:
		return fibonnaciInt(v)
	case float64:
		return fibonnaciFloat(v)
	}
	return nil

}

func fibonnaciUint(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonnaciUint(n-1) + fibonnaciUint(n-2)
}

func fibonnaciInt(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonnaciInt(n-1) + fibonnaciInt(n-2)
}

func fibonnaciFloat(n float64) float64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonnaciFloat(n-1) + fibonnaciFloat(n-2)
}
