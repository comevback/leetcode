package main

import "fmt"

func main() {
	var candies = []int{2, 3, 5, 1, 3}
	var extraCandies = 3
	res := kidsWithCandies(candies, extraCandies)
	fmt.Println(res)
}

// 1.第一次循环找出最大值，第二次循环得到结果
func kidsWithCandies(candies []int, extraCandies int) []bool {
	length := len(candies)
	var result []bool = make([]bool, 0, length)

	// 找出最大值
	var greatest int = 0
	for i := 0; i < length; i++ {
		if candies[i] > greatest {
			greatest = candies[i]
		}
	}

	// 判断加了extraCandies是否可以成为最大值
	for _, v := range candies {
		candiesNum := v + extraCandies
		if candiesNum >= greatest {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	// 返回结果数组
	return result
}
