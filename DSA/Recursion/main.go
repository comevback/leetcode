package main

import "fmt"

func main() {
	res := Fibonacci(6)
	fmt.Println(res)
}

func Factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return Factorial(value-1) * value
	}
}

func Fibonacci(value int) int {
	if value <= 1 {
		return value
	}
	return Fibonacci(value-1) + Fibonacci(value-2)
}
