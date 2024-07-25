# 1475. Final Prices With a Special Discount in a Shop

Solved
Easy
Topics
Companies
Hint
You are given an integer array prices where prices[i] is the price of the ith item in a shop.

There is a special discount for items in the shop. If you buy the ith item, then you will receive a discount equivalent to prices[j] where j is the minimum index such that j > i and prices[j] <= prices[i]. Otherwise, you will not receive any discount at all.

Return an integer array answer where answer[i] is the final price you will pay for the ith item of the shop, considering the special discount.

Example 1:

Input: prices = [8,4,6,2,3]
Output: [4,2,4,2,3]
Explanation:
For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
For items 3 and 4 you will not receive any discount at all.
Example 2:

Input: prices = [1,2,3,4,5]
Output: [1,2,3,4,5]
Explanation: In this case, for all items, you will not receive any discount at all.
Example 3:

Input: prices = [10,1,1,6]
Output: [9,0,1,6]

Constraints:

1 <= prices.length <= 500
1 <= prices[i] <= 1000

---

# Code

```go
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
```
