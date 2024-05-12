package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	hsMap := make(map[string]bool)
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		if hsMap[word] {
			res = append(res, "alreadey_been")
		} else {
			hsMap[word] = true
			res = append(res, word)
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
