# 121. Best Time to Buy and Sell Stock

Easy

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:
> Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
> Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:
> 1 <= prices.length <= 105
0 <= prices[i] <= 104

---

## solution

```go
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

// second try O(n)-O(n)
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
```