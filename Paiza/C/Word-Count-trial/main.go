package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	hsMap := make(map[string]int)
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()

		if _, exist := hsMap[word]; !exist {
			hsMap[word] = 1
			res = append(res, word)
		} else {
			hsMap[word] += 1
		}
	}

	for _, v := range res {
		v = v + " " + strconv.Itoa(hsMap[v])
		fmt.Println(v)
	}
}
