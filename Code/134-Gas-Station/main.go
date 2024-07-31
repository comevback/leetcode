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
