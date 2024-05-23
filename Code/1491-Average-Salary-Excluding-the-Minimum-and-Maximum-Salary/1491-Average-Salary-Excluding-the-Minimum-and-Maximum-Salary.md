# 1491. Average Salary Excluding the Minimum and Maximum Salary

Easy

You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.


Example 1:
> Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500

Example 2:
> Input: salary = [1000,2000,3000]
Output: 2000.00000
Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
 

Constraints:
> 3 <= salary.length <= 100
1000 <= salary[i] <= 106
All the integers of salary are unique.

---

# Code
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	salary := []int{1000, 2000, 3000}
	res := average(salary)
	fmt.Println(res)
}

// 1. 我的解法，一次循环
// 累加每个整数，在循环中找出最大最小值
// 累加结束后，总和减去最大最小值，然后除以总数减二，得到结果
func average(salary []int) float64 {
	var highest, lowest int = math.MinInt, math.MaxInt
	var sum int
	var res float64

	for _, v := range salary {
		sum += v
		if v > highest {
			highest = v
		}
		if v < lowest {
			lowest = v
		}
	}

	sum = sum - highest - lowest
	res = float64(sum) / float64(len(salary)-2)
	return res
}
```