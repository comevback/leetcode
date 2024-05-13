package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	hsMap := make(map[int]bool)
	for i := 1; i <= N; i++ {
		hsMap[i] = true
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if hsMap[input] {
			hsMap[input] = false
		}
	}

	var resCount int = 0
	for _, Value := range hsMap {
		if Value {
			resCount += 1
		}
	}

	fmt.Println(resCount)
}
