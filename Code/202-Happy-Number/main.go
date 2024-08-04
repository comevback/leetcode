package main

import "fmt"

func main() {
	res := isHappy(7)
	fmt.Println(res)
}

// isHappy 判断一个数是否是快乐数
func isHappy(n int) bool {
	// 使用map记录已经计算过的数，避免无限循环
	hsMap := make(map[int]bool)

	// 初始数值
	num := n

	// 当num不等于1时，继续计算
	for num != 1 {
		newNum := 0
		// 计算num的每位数字的平方和
		for num != 0 {
			temp := num % 10      // 获取num的最后一位数字
			newNum += temp * temp // 将该数字的平方加到newNum
			num = num / 10        // 去掉num的最后一位数字
		}

		// 如果newNum已经在map中，说明进入循环，返回false，否则，将newNum加入map中
		if hsMap[newNum] {
			return false
		} else {
			hsMap[newNum] = true
		}
		// 更新num为newNum，继续下一轮计算
		num = newNum
	}

	// 如果num等于1，返回true，说明是快乐数
	return true
}
