package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 0, 0}
	res := findMaxLength(nums)
	fmt.Println(res)
}

// 1.我的解法
func findMaxLength1(nums []int) int {
	maxLength := 0
	zerofix := make([]int, len(nums)+1)
	onefix := make([]int, len(nums)+1)
	zeroSum, oneSum := 0, 0

	for i := 0; i < len(nums)+1; i++ {
		zerofix[i] = zeroSum
		onefix[i] = oneSum
		if i < len(nums) {
			if nums[i] == 0 {
				zeroSum += 1
			} else {
				oneSum += 1
			}
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if zerofix[j+1]-zerofix[i] == onefix[j+1]-onefix[i] {
				if j-i+1 > maxLength {
					maxLength = j - i + 1
				}
			}
		}
	}

	return maxLength
}

// =======================================================================================================
// 2.前缀 + 哈希表
func findMaxLength(nums []int) int {
	maxLength := 0                    // 记录最长的长度
	prefix := make([]int, len(nums))  // 用于存储前缀和，前面每有一个0减1，每有一个1加1，包括自身
	preMap := make(map[int][]int)     // 用于存储前缀和的位置
	preMap[0] = append(preMap[0], -1) // 初始化，当前缀和为0时，位置设为-1，便于计算长度
	sum := 0                          // 记录当前的前缀和

	// 遍历数组，计算前缀和
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum -= 1 // 遇到0，将sum减1
		} else {
			sum += 1 // 遇到1，将sum加1
		}
		prefix[i] = sum                      // 存储当前的前缀和
		preMap[sum] = append(preMap[sum], i) // 将当前的前缀和位置存入哈希表
	}

	// 遍历哈希表，计算最长的子数组长度
	for i := range preMap {
		// 计算当前前缀和对应的最长子数组长度，就是拥有同一个前缀值的最后和最前的索引相减
		length := preMap[i][len(preMap[i])-1] - preMap[i][0]
		if length > maxLength {
			maxLength = length // 更新最大长度
		}
	}

	return maxLength // 返回最长子数组长度
}

// ========================================================================================================
// 思想和我的解法一样，但只遍历一次，在遍历过程中记录最大长度
func findMaxLength2(nums []int) int {
	hashcount := make(map[int]int) // 哈希表，用于存储每个差值第一次出现的位置
	ans, ones, zeros := 0, 0, 0    // 初始化变量

	for i, val := range nums { // 遍历数组
		if val == 1 {
			ones++ // 遇到 1，增加 ones 计数
		} else {
			zeros++ // 遇到 0，增加 zeros 计数
		}

		// 检查哈希表中是否已有当前的差值
		if _, ok := hashcount[ones-zeros]; !ok {
			hashcount[ones-zeros] = i // 记录当前差值第一次出现的位置
		}
		// 如果 ones 和 zeros 相等，更新 ans 为当前长度
		if ones == zeros {
			ans = ones + zeros
		} else {
			// 否则，找到之前相同差值的位置，计算子数组长度，并更新 ans
			idx := hashcount[ones-zeros]
			ans = max(ans, i-idx)
		}
	}

	return ans // 返回最长子数组长度
}

// 辅助函数，用于计算两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
