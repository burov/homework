# Task

* Create go module `computator`
* Create package `fibonacci`
* Create type specific function for uint, int, float64
* Create general function that accept and return an empty interface{} [What is the empty interfaces?](https://tour.golang.org/methods/14)
* In the general function using type switch call proper function depending on the type of the function parameter [What is the type switch?](https://tour.golang.org/methods/16)
* In computator module create file `main.go`, implement `func main() {...}` to run application from console [How to read number from the terminal](https://golang.org/pkg/fmt/#Scanf)
* In the main function read from the console user input, if the input not a number - return an error, if it is the number find most specific type from the list (uint, int, float64), convert input to the specific type and call `fibonacci` function 

Functions signature should be equals to functions presented below 
``` go
func fibonacci(n interface{}) (interface{}) {...}

func fibonnaciUint(n uint) (uint) {...}

func fibonnaciInt(n int) (int) {...}

func fibonnaciFloat(n float64) {...}
```
