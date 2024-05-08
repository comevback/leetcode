package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var times int = 2
	var preStr []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		preStr = append(preStr, scanner.Text())
		times -= 1
		if times < 1 {
			break
		}
	}

	if preStr[0] == preStr[1] {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
