package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	length := 12
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	sumDis := 0

	for scanner.Scan() {
		var num int
		var add int
		char := scanner.Text()
		if !(char <= "9" && char >= "0") {
			continue
		}
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}

		if num == 0 {
			add = 12
		} else {
			add = num + 2
		}

		sumDis += add * 2

		length -= 1
		if length < 1 {
			break
		}
	}

	fmt.Println(sumDis)
}
