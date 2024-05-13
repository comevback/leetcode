package main

import (
	"fmt"
)

func main() {
	var N, X, Y int
	fmt.Scanf("%d %d %d", &N, &X, &Y)
	res := []string{}

	for i := 1; i <= N; i++ {
		if i%X == 0 && i%Y == 0 {
			res = append(res, "AB")
		} else if i%X == 0 {
			res = append(res, "A")
		} else if i%Y == 0 {
			res = append(res, "B")
		} else {
			res = append(res, "N")
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
