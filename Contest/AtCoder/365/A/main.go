package main

import "fmt"

func main() {
	var year int
	fmt.Scan(&year)
	res := 0

	if year%4 != 0 {
		res = 365
	}

	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			res = 365
		} else {
			res = 366
		}
	}

	fmt.Println(res)
}
