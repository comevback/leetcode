package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nums int
	fmt.Scan(&nums)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	numArr := strings.Split(str, "")
	win := 0
	last := ""

	dp := make([][]int, len(numArr))
	for i := range dp {

	}

	// R : 0
	// P : 1
	// S : 2

	for i, v := range numArr {
		switch v {
		case "R":
			if last != "P" {
				last = "P"
				win += 1
			} else {
				last = "R"
			}
		case "P":
			if last != "S" {
				last = "S"
				win += 1
			} else {
				last = "P"
			}
		case "S":
			if last != "R" {
				last = "R"
				win += 1
			} else {
				last = "S"
			}
		}
	}

	fmt.Println(win)
}
