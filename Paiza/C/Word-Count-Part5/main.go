package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	hsMap := make(map[string]int)
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		if hsMap[word] >= 1 {
			hsMap[word] += 1
		} else {
			hsMap[word] = 1
			res = append(res, word)
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(1)
	}
}
