package main

import "fmt"

func main() {
	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices1(prices)
	fmt.Println(res)
}

// finalPrices 通过一个栈来处理每个价格后面的可能折扣。
// finalPrices processes possible discounts for each price using a stack.
func finalPrices(prices []int) []int {
	res := make([]int, len(prices)) // 初始化结果数组
	temp := []int{}                 // 使用 temp 栈来跟踪之前的价格

	for i := len(prices) - 1; i >= 0; i-- {
		dis := 0 // 初始化折扣为0
		j := len(temp) - 1
		// 寻找第一个小于等于当前价格的价格作为折扣
		for j >= 0 && prices[i] < temp[j] {
			j -= 1
		}
		if j >= 0 {
			dis = temp[j] // 设置折扣
		}

		res[i] = prices[i] - dis       // 计算最终价格
		temp = append(temp, prices[i]) // 将当前价格压入栈
	}
	return res
}

// **********************************************  先求得下一个小于等于本身的树  ********************************************

// finalPrices1 通过先找出每个价格的下一个更小或相等的价格，然后计算最终价格。
// finalPrices1 first finds the next smaller or equal price for each price, then calculates the final prices.
func finalPrices1(prices []int) []int {
	res := make([]int, len(prices))
	dis := findNextSmaller(prices) // 找到每个价格后面第一个更小或相等的价格

	for i := 0; i < len(dis); i++ {
		res[i] = prices[i] - dis[i] // 计算最终价格
	}
	return res
}

// findNextSmaller 寻找每个价格后面第一个更小或相等的价格。
// findNextSmaller finds the next smaller or equal price for each price.
func findNextSmaller(prices []int) []int {
	res := make([]int, len(prices))
	temp := []int{} // 使用栈来跟踪价格

	for i := len(prices) - 1; i >= 0; i-- {
		// 移除所有大于当前价格的元素
		for len(temp) > 0 && prices[i] < temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		}

		if len(temp) == 0 {
			res[i] = 0 // 如果没有找到合适的折扣价格，结果为0
		} else {
			res[i] = temp[len(temp)-1] // 使用找到的折扣价格
		}
		temp = append(temp, prices[i]) // 将当前价格压入栈
	}
	return res
}
