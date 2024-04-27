package main

import "fmt"

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(prices))
}

// first try 双循环，时间复杂度O(n^2),不好
func maxProfit(prices []int) int {
	length := len(prices)
	maxP := 0
	for i, v := range prices {
		for j := i + 1; j < length; j++ {
			if prices[j]-v > maxP {
				maxP = prices[j] - v
			}
		}
	}
	return maxP
}

//========================================================================================================================

// 把每一天和前一天相差的差值，统计为一个数组 O(n)-O(n)
func maxProfit2(prices []int) int {
	length := len(prices)

	// 把每一天和前一天相差的差值，统计为一个数组
	var abs []int
	for i := 1; i < length; i++ {
		abs = append(abs, prices[i]-prices[i-1])
		fmt.Println(abs)
	}

	// 设一个profile，如果到下一天，累加的利益大于零，就继续下去，如果小于零，就归零，从这一天开始，算第一天，并且统计曾经达到过的最大值。
	var profit int = 0
	var maxP int = 0
	for _, v := range abs {
		if profit+v > 0 {
			profit += v
		} else {
			profit = 0
		}

		if profit > maxP {
			maxP = profit
		}
	}

	// 返回达到过的最大值
	return maxP
}

//========================================================================================================================

// other solution
func maxProfit3(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit
}

// or
func maxProfit4(prices []int) int {
	var maxProfit = 0
	var minPrice = prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

//========================================================================================================================
