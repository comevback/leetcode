package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var times int
	fmt.Scan(&times)
	scanner := bufio.NewScanner(os.Stdin)
	strikeCount := 0
	ballCount := 0
	res := []string{}

	for scanner.Scan() {
		str := scanner.Text()
		if str == "strike" {
			strikeCount += 1
			if strikeCount == 3 {
				res = append(res, "out!")
				break
			} else {
				res = append(res, "strike!")
			}
		}

		if str == "ball" {
			ballCount += 1
			if ballCount == 4 {
				res = append(res, "fourball!")
				break
			} else {
				res = append(res, "ball!")
			}
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
