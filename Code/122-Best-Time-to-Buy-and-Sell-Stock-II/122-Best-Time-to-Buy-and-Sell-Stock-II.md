# 122. Best Time to Buy and Sell Stock II

Medium

You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

 

Example 1:
> Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

Example 2:
> Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

Example 3:
> Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
 

Constraints:
> 1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104

```go
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
```