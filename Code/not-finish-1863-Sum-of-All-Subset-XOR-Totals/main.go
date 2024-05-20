package main

func main() {

}

func subsetXORSum(nums []int) int {
	sum := 0
	res := getSum(nums, 0, sum)
	return res
}

// 递归函数
func getSum(nums []int, index int, sumXOR int) int {
	// 一直递归到长度为len(nums)，返回一路上获得的sumXOR
	if index == len(nums) {
		return sumXOR
	}

	// **** 关键 ****  如果当前长度还不到len(nums)，继续递归，
	return getSum(nums, index+1, sumXOR) + getSum(nums, index+1, sumXOR^nums[index])
}

// 递归函数的工作原理
// 该行代码：

// return getSum(nums, index+1, sumXOR) + getSum(nums, index+1, sumXOR^nums[index])
// 表达了两种可能的递归调用：

// 不包含当前索引元素：getSum(nums, index+1, sumXOR)
// 包含当前索引元素：getSum(nums, index+1, sumXOR^nums[index])
// 这两个调用覆盖了所有包含或不包含当前索引元素的子集情况。我们逐步遍历这个过程。

// 示例：nums = [1, 2]
// 初始调用
// 调用 getSum(nums, 0, 0) 开始，此时 index = 0，sumXOR = 0。
// 第一层递归
// 不包含元素 1：调用 getSum(nums, 1, 0)
// 包含元素 1：调用 getSum(nums, 1, 0^1)，这里 0^1 = 1
// 第二层递归（对于第一层的每个调用）
// 对于 getSum(nums, 1, 0)：

// 不包含元素 2：调用 getSum(nums, 2, 0)，到达基础情况，返回 0。
// 包含元素 2：调用 getSum(nums, 2, 0^2)，这里 0^2 = 2，到达基础情况，返回 2。
// 对于 getSum(nums, 1, 1)：

// 不包含元素 2：调用 getSum(nums, 2, 1)，到达基础情况，返回 1。
// 包含元素 2：调用 getSum(nums, 2, 1^2)，这里 1^2 = 3，到达基础情况，返回 3。
// 汇总结果
// 将所有返回的值相加：

// getSum(nums, 2, 0) -> 0
// getSum(nums, 2, 2) -> 2
// getSum(nums, 2, 1) -> 1
// getSum(nums, 2, 3) -> 3
// 所以，总和是 0 + 2 + 1 + 3 = 6。
