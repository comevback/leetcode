package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line int
	fmt.Scan(&line)

	scanner := bufio.NewScanner(os.Stdin)
	var str string

	for scanner.Scan() {
		text := scanner.Text()
		str += text + "\n"
		line -= 1
		if line < 1 {
			break
		}
	}

	fmt.Println(str)
}
