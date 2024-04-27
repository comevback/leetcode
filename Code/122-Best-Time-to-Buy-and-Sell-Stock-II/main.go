package main

import "fmt"

func main() {
	var arr []int = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(arr))
}

// ===========================================================================================================================
// 贪心算法

// 1. 同121，把每一天和前一天相差的差值，统计为一个数组，把所有正数相加，即为最大利益
func maxProfit(prices []int) int {
	// 如果只有一个数在数组里，没有利益，返回0
	length := len(prices)
	if length < 2 {
		return 0
	}
	// 生成差值数组
	var profit []int = make([]int, 0, length-1)
	for i := 1; i < length; i++ {
		profit = append(profit, prices[i]-prices[i-1])
	}

	// 计算最大利润，只要是正数，就加起来
	fmt.Println(profit)
	var maxP int = 0
	for _, v := range profit {
		if v > 0 {
			maxP += v
		}
	}
	// 返回最大利润
	return maxP
}

// 2. 差值数组的简化版
func maxProfit1(prices []int) int {
	// 如果只有一个数在数组里，没有利益，返回0
	length := len(prices)
	if length < 2 {
		return 0
	}
	// 只要差值是正数，就加到最大利益里
	var maxP int = 0
	for i := 1; i < length; i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			maxP += change
		}
	}
	// 返回最大利益
	return maxP
}

// ===========================================================================================================================

// 2.动态规划 O(n)-O(n)
func maxProfit2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	// 生成一个二维数组，第一列是不持有股票的利润，第二列是持有股票的利润
	// 例如：profit[0][0] = 0, profit[0][1] = -prices[0]
	var profit [][2]int = make([][2]int, length)
	profit[0][0] = 0
	profit[0][1] = -prices[0]

	// 从第二天开始，根据前一天的利润，计算今天的利润，直到最后一天
	for i := 1; i < length; i++ {
		profit[i][0] = max(profit[i-1][0], profit[i-1][1]+prices[i])
		profit[i][1] = max(profit[i-1][1], profit[i-1][0]-prices[i])
	}

	// 返回最后一天不持有股票的利润
	return profit[length-1][0]
}

// 动态规划，优化空间复杂度 O(n)-O(1)
func maxProfit3(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	prevNoStock := 0
	prevStock := -prices[0]

	for i := 1; i < length; i++ {
		currentNoStock := max(prevNoStock, prevStock+prices[i])
		currentStock := max(prevStock, prevNoStock-prices[i])
		prevNoStock = currentNoStock
		prevStock = currentStock
	}

	return prevNoStock
}

// 计算两个数的最大值
func max(m int, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}
