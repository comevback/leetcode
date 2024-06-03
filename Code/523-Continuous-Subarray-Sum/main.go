package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 0, 4, 0, 4, 0}
	res := checkSubarraySum(nums, 5)
	fmt.Println(res)
}

// 检查是否存在和为 k 的连续子数组的函数
func checkSubarraySum(nums []int, k int) bool {
	// 如果数组长度小于 2，返回 false，因为至少需要两个元素才能形成子数组
	if len(nums) < 2 {
		return false
	}

	// 前缀和数组，长度比 nums 长 1，用于存储前缀和
	prefix := make([]int, len(nums)+1)
	sum := 0
	// 余数映射，用于存储前缀和对 k 取模后的结果及其对应的索引
	divideMap := make(map[int][]int)

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 如果当前元素为 0，且下一个元素也是 0，则返回 true（存在长度为 2 的子数组和为 0）
		if nums[i] == 0 {
			if i < len(nums)-1 && nums[i+1] == 0 {
				return true
			}
		}
		sum += nums[i]    // 更新当前前缀和
		prefix[i+1] = sum // 存储当前前缀和到前缀和数组
		// 如果当前前缀和对 k 取模为 0，则返回 true（存在长度至少为 2 的子数组和为 k 的倍数）
		if i > 0 && prefix[i+1]%k == 0 {
			return true
		}
		// 将当前前缀和对 k 取模的结果及其索引存入映射
		divideMap[prefix[i+1]%k] = append(divideMap[prefix[i+1]%k], i)
		// 检查是否存在两个索引，它们之间的子数组长度至少为 2 且前缀和的差对 k 取模为 0
		if divideMap[prefix[i+1]%k][len(divideMap[prefix[i+1]%k])-1]-divideMap[prefix[i+1]%k][0] >= 2 {
			return true
		}
	}

	return false // 没有找到满足条件的子数组，返回 false
}
