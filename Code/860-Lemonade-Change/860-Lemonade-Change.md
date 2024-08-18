# 860. Lemonade Change

Solved
Easy
Topics
Companies
At a lemonade stand, each lemonade costs $5. Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. You must provide the correct change to each customer so that the net transaction is that the customer pays $5.

Note that you do not have any change in hand at first.

Given an integer array bills where bills[i] is the bill the ith customer pays, return true if you can provide every customer with the correct change, or false otherwise.

Example 1:

Input: bills = [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.
Example 2:

Input: bills = [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can not give the change of $15 back because we only have two $10 bills.
Since not every customer received the correct change, the answer is false.

Constraints:

1 <= bills.length <= 105
bills[i] is either 5, 10, or 20.

---

# Code

```go
package main

import "fmt"

func main() {
	// 示例输入：顾客支付的钞票序列
	bills := []int{5, 5, 5, 10, 20}
	// 调用函数，判断是否可以正确找零
	res := lemonadeChange(bills)
	// 输出结果
	fmt.Println(res)
}

// lemonadeChange 函数判断能否给每个顾客正确找零
func lemonadeChange(bills []int) bool {
	// have5 和 have10 分别记录当前持有的 5 美元和 10 美元的数量
	have5 := 0
	have10 := 0

	// 遍历所有顾客支付的钞票
	for _, v := range bills {
		if v == 5 {
			// 如果顾客支付 5 美元，直接收下，不需要找零
			have5 += 1
		} else if v == 10 {
			// 如果顾客支付 10 美元，需要找零 5 美元
			if have5 < 1 {
				// 如果没有 5 美元，无法找零，返回 false
				return false
			}
			// 找零成功，减少 5 美元的数量，增加 10 美元的数量
			have5 -= 1
			have10 += 1
		} else if v == 20 {
			// 如果顾客支付 20 美元，需要找零 15 美元
			if have10 > 0 && have5 > 0 {
				// 优先使用一个 10 美元和一个 5 美元找零
				have10 -= 1
				have5 -= 1
			} else if have5 > 2 {
				// 否则，如果有至少 3 张 5 美元的钞票，用它们找零
				have5 -= 3
			} else {
				// 如果既不能用 10 美元+5 美元，也不能用 3 张 5 美元找零，返回 false
				return false
			}
		}
	}

	// 如果所有顾客都能正确找零，返回 true
	return true
}
```
