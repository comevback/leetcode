package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		res = append(res, word)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
