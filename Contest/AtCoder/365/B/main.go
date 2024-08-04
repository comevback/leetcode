package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var times int
	fmt.Scan(&times)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	max, sec := -1, -1
	maxIndex, secIndex := -1, -1
	index := 1

	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)

		if num > max {
			sec = max
			secIndex = maxIndex
			max = num
			maxIndex = index
		} else if num > sec {
			sec = num
			secIndex = index
		}
		index += 1

		times -= 1
		if times < 1 {
			break
		}
	}

	fmt.Println(secIndex)
}
