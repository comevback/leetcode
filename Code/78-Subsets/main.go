package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	res := subsets(arr)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	getSubs(nums, &res, []int{}, 0)
	return res
}

// 这个程序的思想在于，每一次递归，当前的 nums[index] 到底加不加入temp，都考虑两种情况
// index从0开始，到len(temp)-1
// 1. nums[0]，是否加入temp，分化为两支，
// 2. nums[1], 是否加入，分化为四支，
// 3. nums[3], 是否加入，分化为八支......
// 当每一支的index都到len(temp)之后，不能再加了，所以把这一支的切片，加入到结果数组中。
func getSubs(nums []int, res *[][]int, temp []int, index int) {
	if index == len(nums) {
		t := make([]int, len(temp))
		copy(t, temp)
		*res = append(*res, t)
	} else {
		getSubs(nums, res, temp, index+1)
		temp = append(temp, nums[index])
		getSubs(nums, res, temp, index+1)
	}
}
