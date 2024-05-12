package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	res := "No"

	for scanner.Scan() {
		word := scanner.Text()
		if word == "red" {
			res = "Yes"
		}
	}

	fmt.Println(res)
}
