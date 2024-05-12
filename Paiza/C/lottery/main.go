package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var priceNumber, n int
	fmt.Scan(&priceNumber)
	fmt.Scan(&n)
	strPrice := strconv.Itoa(priceNumber)

	var res []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()

		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		if number == priceNumber {
			res = append(res, "first")
		} else if (number == priceNumber+1 || number == priceNumber-1) && number > 100000 && number < 199999 {
			res = append(res, "adjacent")
		} else if str[2:] == strPrice[2:] {
			res = append(res, "second")
		} else if str[3:] == strPrice[3:] {
			res = append(res, "third")
		} else {
			res = append(res, "blank")
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
