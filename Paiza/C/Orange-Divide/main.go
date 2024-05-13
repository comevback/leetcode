package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, num int
	fmt.Scanf("%d %d", &N, &num)
	res := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		result := findNearestValue(N, weight)
		res = append(res, result)

		num -= 1
		if num < 1 {
			break
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

// 找最接近箱子功能
func findNearestValue(N int, weight int) int {
	if weight <= N {
		return N
	}

	min := N - weight
	if min < 0 {
		min = -min
	}

	times := 1
	for {
		temp := N*times - weight
		if temp < 0 {
			temp = -temp
		}
		if temp <= min {
			min = temp
			times += 1
		} else {
			return N * (times - 1)
		}
	}
}
