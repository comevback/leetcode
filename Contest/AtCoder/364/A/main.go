package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var times int
	fmt.Scan(&times)
	lastSweet := false
	res := true

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "salty" {
			lastSweet = false
		} else if str == "sweet" {
			if lastSweet {
				res = false
			}
			lastSweet = true
		}

		times -= 1
		if times < 1 {
			break
		}
	}

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
