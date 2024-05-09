# 3075. Maximize Happiness of Selected Children（Max Heap / Priority Queue）

Medium

Hint
You are given an array happiness of length n, and a positive integer k.

There are n children standing in a queue, where the ith child has happiness value happiness[i]. You want to select k children from these n children in k turns.

In each turn, when you select a child, the happiness value of all the children that have not been selected till now decreases by 1. Note that the happiness value cannot become negative and gets decremented only if it is positive.

Return the maximum sum of the happiness values of the selected children you can achieve by selecting k children.

 

Example 1:
> Input: happiness = [1,2,3], k = 2
Output: 4
Explanation: We can pick 2 children in the following way:
- Pick the child with the happiness value == 3. The happiness value of the remaining children becomes [0,1].
- Pick the child with the happiness value == 1. The happiness value of the remaining child becomes [0]. Note that the happiness value cannot become less than 0.
The sum of the happiness values of the selected children is 3 + 1 = 4.

Example 2:
> Input: happiness = [1,1,1,1], k = 2
Output: 1
Explanation: We can pick 2 children in the following way:
- Pick any child with the happiness value == 1. The happiness value of the remaining children becomes [0,0,0].
- Pick the child with the happiness value == 0. The happiness value of the remaining child becomes [0,0].
The sum of the happiness values of the selected children is 1 + 0 = 1.

Example 3:
> Input: happiness = [2,3,4,5], k = 1
Output: 5
Explanation: We can pick 1 child in the following way:
- Pick the child with the happiness value == 5. The happiness value of the remaining children becomes [1,2,3].
The sum of the happiness values of the selected children is 5.
 

Constraints:
> 1 <= n == happiness.length <= 2 * 105
1 <= happiness[i] <= 108
1 <= k <= n

---

# 代码

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var happiness []int = []int{1, 2, 3}
	res := maximumHappinessSum2(happiness, 2)
	fmt.Println(res)
}

// ===============================================  排序后直接操作，双循环 O(n^2) 错误  ================================================
func maximumHappinessSum1(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var sum int64 = 0
	for i := k; i > 0; i-- {
		if len(happiness) != 0 {
			sum += int64(happiness[len(happiness)-1])
			happiness = happiness[:len(happiness)-1]
			for i := 0; i < len(happiness); i++ {
				if happiness[i] > 0 {
					happiness[i] -= 1
				}
			}
		}
	}

	return sum
}

// ==============================================  排序后求和时减(次数-1) N(NlogN) 可以 ===================================================
func maximumHappinessSum(happiness []int, k int) int64 {
	// 排序
	sort.Ints(happiness)
	// 定义累加的最大值
	var sum int64 = 0
	// 每次累加减去的minus
	minus := 0
	// 最多循环k次
	for i := k; i > 0; i-- {
		// 如果数组不为空，继续循环
		if len(happiness) != 0 {
			// 如果数组中最大的数减去minus大于0，累加这个数-minus，数组去掉尾部最大值，minus加一
			if happiness[len(happiness)-1]-minus > 0 {
				sum += int64(happiness[len(happiness)-1] - minus)
				happiness = happiness[:len(happiness)-1]
				minus += 1
			} else {
				// 如果数组最大值减minus小于等于0，则数组中所有值减minus都小于等于0，没必要再累加，跳出循环
				break
			}
		}
	}
	// 返回累加的最大值
	return sum
}

// ===================================================  每次查找最大值 时间复杂度太高  =================================================
func maximumHappinessSum2(happiness []int, k int) int64 {
	sum := 0

	for i := 0; i < k; i++ {
		if len(happiness) != 0 {
			add := findandDeleteMax(&happiness) - i
			if add > 0 {
				sum += add
			} else {
				break
			}
		} else {
			break
		}
	}

	return int64(sum)
}

func findandDeleteMax(nums *[]int) int {
	max := math.MinInt
	var maxIndex int = -1

	for i, v := range *nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}

	if maxIndex != -1 && len(*nums) > 0 {
		if maxIndex == len(*nums)-1 {
			*nums = (*nums)[:len(*nums)-1] // 如果最大值在最后，直接截断末尾元素
		} else {
			*nums = append((*nums)[:maxIndex], (*nums)[maxIndex+1:]...) // 正确删除 maxIndex 处的元素
		}
	}

	return max
}

// ==============================================　Max Heap / Priority Queue + Greedy　=============================================
func maximumHappinessSum3(happiness []int, k int) int64 {
	// use Max Heap / Priority Queue
	return int64(0)
}
```