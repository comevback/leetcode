package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var nums, maxMoney int
	fmt.Scanf("%d %d", &nums, &maxMoney)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	expens := []int{}
	sum := 0
	max := math.MinInt

	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		if num > max {
			max = num
		}
		sum += num
		expens = append(expens, num)

		nums -= 1
		if nums < 1 {
			break
		}
	}

	if sum <= maxMoney {
		fmt.Println("infinite")
		return
	}

	diff := sum - maxMoney

	left, right := 0, max
	for left < right {
		temp := 0
		mid := left + (right-left)/2
		for _, v := range expens {
			temp += v - mid
		}
		if diff <= temp {
			fmt.Println()
			return
		}
	}
}
