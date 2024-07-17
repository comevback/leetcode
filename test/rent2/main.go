package main

import "fmt"

func main() {
	res := fact(5) // テストの例
	fmt.Println(res)
}

func fact(n int) int {
	if n < 0 {
		fmt.Println("invaild input")
		return -1
	}

	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
