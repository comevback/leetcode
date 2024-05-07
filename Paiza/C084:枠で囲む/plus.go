package main

import "fmt"

func main() {
	var inputed string
	fmt.Scan(&inputed)
	resStr := processString(inputed)
	fmt.Println(resStr)
}

func processString(str string) string {
	newStr := "+" + str + "+"
	var plusLine string
	for i := 0; i < len(newStr); i++ {
		plusLine += "+"
	}

	newStr = plusLine + "\n" + newStr + "\n" + plusLine
	return newStr
}
