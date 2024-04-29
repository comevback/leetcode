package main

import "fmt"

func main() {
	var nums = []int{5, 4, -1, 7, 8}
	res := maxSubArray1(nums)
	fmt.Println(res)
}

// ================================================================================================================================
// 1.贪心算法 O(n)-O(1)
// 初始最大值为第一个元素的值，如果在之前的累加中sum已经小于0，就把之前的抛掉
// 在遍历累加过程中，sum达到的最大的值就是maxSum，返回maxSum
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var sum int
	var maxSum int = nums[0]
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

// ================================================================================================================================
// 2.动态规划，Kadane算法
// 计算以第i个数结尾的最大字数组之和, dp[i] = max(nums[i], dp[i-1]+nums[i])
// 遍历数组，找到dp[i]的最大值
func maxSubArray1(nums []int) int {
	maxSum := nums[0]
	var dp []int = make([]int, len(nums))
	dp[0] = nums[0]
	// 在遍历累加过程中，dp[i]达到的最大的值就是maxSum，返回maxSum
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	fmt.Print(dp)
	return maxSum
}

// ================================================================================================================================
// 3.分治法
// 步骤：
// 1.分解：将数组分成两半。
// 2.解决：递归地找出左半部和右半部的最大子数组和。
// 3.合并：最大子数组可能跨越了两个子数组的中间位置，所以还需要计算跨中点的最大子数组和，并与前两步的结果比较。
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	leftMax := maxSubArray(nums[:mid])
	rightMax := maxSubArray(nums[mid:])

	midMax := findMaxCrossingSum(nums, mid)

	return max(max(leftMax, rightMax), midMax)
}

func findMaxCrossingSum(nums []int, mid int) int {
	// 这是将leftSum和rightSum初始化int可达到的最小值
	leftSum := -1 << 31
	rightSum := -1 << 31

	// 从中间开始向左累加，能达到的包含中点-1的字数组和最大值leftSum
	sum := 0
	for i := mid - 1; i >= 0; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	// 从中间开始向右累加，能达到的包含中点-1的字数组和最大值rightSum
	sum = 0
	for i := mid; i < len(nums); i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// 返回包含中点的左右字数组最大之和，也就是跨中点的最大子数组和
	return leftSum + rightSum
}
