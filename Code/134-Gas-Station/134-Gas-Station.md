# 134. Gas Station

Solved
Medium
Topics
Companies
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

Example 1:

Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
Example 2:

Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.

Constraints:

n == gas.length == cost.length
1 <= n <= 105
0 <= gas[i], cost[i] <= 104

---

# Code

```go
package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res) // 输出可以完成环形路线的起始加油站索引
}

// canCompleteCircuit1 通过暴力解法尝试找出可以完成一圈的起始加油站
// 使用双循环方式，检查每个加油站是否可以作为起点完成一圈
func canCompleteCircuit1(gas []int, cost []int) int {
	length := len(gas) // 获取加油站的数量
	res := -1          // 默认没有找到可以完成一圈的起始加油站

	// 外层循环遍历每一个加油站作为起点
	for i := 0; i < length; i++ {
		if gas[i] < cost[i] {
			// 如果当前加油站的油不足以到达下一个加油站，直接跳过
			continue
		} else {
			amount := 0 // 从当前加油站开始的油量余额
			j := i      // 从当前加油站出发
			res = i     // 假设从当前加油站可以完成一圈
			// 内层循环遍历完成一圈的过程
			for j < i+length {
				// 计算从当前加油站到下一个加油站后的油量余额
				amount += gas[j%length]
				amount -= cost[j%length]
				j += 1
				// 如果油量余额小于0，说明无法从这个加油站开始完成一圈
				if amount < 0 {
					res = -1
					break
				}
			}
			// 如果从当前加油站出发可以完成一圈，则停止搜索
			if res != -1 {
				break
			}
		}
	}

	return res // 返回结果
}

// -----------------------------------------------------------------------------------------------------------
// canCompleteCircuit 通过一次遍历确定是否有可能从某个加油站开始完成一圈
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0     // 总的油量减去总的花费，用来判断整个路线是否可行
	tempSum := 0 // 从某个加油站开始计算的油量余额
	res := -1    // 默认没有找到可以完成一圈的起始加油站

	// 遍历每一个加油站
	for i := 0; i < len(gas); i++ {
		temp := gas[i] - cost[i] // 计算当前加油站的油量和花费的差值
		// 判断从当前或之前某个加油站累积的油量余额是否足够到达下一个加油站
		if tempSum+temp >= 0 {
			if res == -1 {
				res = i // 如果是第一次油量余额足够，设置当前加油站为起点
			}
			tempSum += temp // 累加油量余额
		} else {
			res = -1    // 如果油量余额不足，重置起始加油站
			tempSum = 0 // 重置油量余额
		}
		sum += temp // 累计总油量差
	}
	// 最后检查总的油量是否足够
	if sum < 0 {
		return -1 // 如果总的油量不足，返回-1
	}

	return res // 返回可以完成一圈的起始加油站索引
}
```
