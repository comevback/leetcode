package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func paiza() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a word:")

	scanner.Scan()
	str := scanner.Text()
	str = strings.TrimSpace(str)

	if str == "paiza" {
		fmt.Println("YES")
	} else {
		fmt.Println("No")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
