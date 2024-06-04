# 1094. Car Pooling
Solved
Medium
Topics
Companies
Hint
There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).

You are given the integer capacity and an array trips where trips[i] = [numPassengersi, fromi, toi] indicates that the ith trip has numPassengersi passengers and the locations to pick them up and drop them off are fromi and toi respectively. The locations are given as the number of kilometers due east from the car's initial location.

Return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.


Example 1:
> Input: trips = [[2,1,5],[3,3,7]], capacity = 4
Output: false

Example 2:
> Input: trips = [[2,1,5],[3,3,7]], capacity = 5
Output: true

Constraints:

1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi <= 100
0 <= fromi < toi <= 1000
1 <= capacity <= 105

---

# Code
```go
package main

import "fmt"

func main() {
	trips := [][]int{{9, 0, 1}, {3, 3, 7}}
	res := carPooling(trips, 4)
	fmt.Println(res)
}

func carPooling(trips [][]int, capacity int) bool {
	// 初始化差分数组，用于记录每个站点的乘客变化
	diff := []int{}

	// 遍历每个行程记录
	for _, v := range trips {
		// 如果当前行程的下车站点超过当前差分数组的长度，扩展差分数组
		if v[2] > len(diff)-1 {
			diff = append(diff, make([]int, v[2]+1-len(diff))...)
		}

		// 在上车站点增加乘客数
		diff[v[1]] += v[0]
		// 在下车站点减少乘客数
		diff[v[2]] -= v[0]
	}

	// 初始化累积乘客数组，用于计算每个站点的总乘客数
	res := make([]int, len(diff))
	// 第一个站点的乘客数直接等于差分数组的第一个值
	res[0] = diff[0]
	// 如果第一个站点的乘客数超过容量，返回false
	if res[0] > capacity {
		return false
	}
	// 计算每个站点的累积乘客数，并检查是否超过容量
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
		if res[i] > capacity {
			return false
		}
	}

	// 如果所有站点的乘客数都不超过容量，返回true
	return true
}
```