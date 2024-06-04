# 1109. Corporate Flight Bookings
Solved
Medium
Topics
Companies
There are n flights that are labeled from 1 to n.

You are given an array of flight bookings bookings, where bookings[i] = [firsti, lasti, seatsi] represents a booking for flights firsti through lasti (inclusive) with seatsi seats reserved for each flight in the range.

Return an array answer of length n, where answer[i] is the total number of seats reserved for flight i.

Example 1:
> Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
Output: [10,55,45,25,25]
Explanation:
Flight labels:        1   2   3   4   5
Booking 1 reserved:  10  10
Booking 2 reserved:      20  20
Booking 3 reserved:      25  25  25  25
Total seats:         10  55  45  25  25
Hence, answer = [10,55,45,25,25]

Example 2:
> Input: bookings = [[1,2,10],[2,2,15]], n = 2
Output: [10,25]
Explanation:
Flight labels:        1   2
Booking 1 reserved:  10  10
Booking 2 reserved:      15
Total seats:         10  25
Hence, answer = [10,25]

 

Constraints:

1 <= n <= 2 * 104
1 <= bookings.length <= 2 * 104
bookings[i].length == 3
1 <= firsti <= lasti <= n
1 <= seatsi <= 104

---

# Code
```go
package main

import "fmt"

func main() {
	res := corpFlightBookings([][]int{{3, 3, 5}, {1, 3, 20}, {1, 2, 15}}, 3)
	fmt.Println(res)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	// 初始化一个长度为n的差分数组
	diff := make([]int, n)

	// 遍历每个预订记录，并更新差分数组
	for _, v := range bookings {
		// 将预订人数加到起始位置（v[0]-1）
		diff[v[0]-1] += v[2]
		// 如果结束位置在数组范围内，将预订人数从结束位置+1处减去
		if v[1] < len(diff) {
			diff[v[1]] -= v[2]
		}
	}

	// 初始化结果数组
	res := make([]int, len(diff))
	// 第一个位置直接等于差分数组的第一个位置
	res[0] = diff[0]
	// 通过累加差分数组计算原数组的每个位置的值
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}

	// 返回结果数组
	return res
}
```