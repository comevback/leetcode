package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(runningSum(arr1))
}

// func runningSum(nums []int) []int {
// 	// 定义一个整数型切片，长度为0，容量为给出的nums的长度len(nums)
// 	var res []int = make([]int, 0, len(nums))

// 	// 定义一个整数，存储目前累加的值
// 	var tem int = 0

// 	// 循环nums切片，累加tem，每次的tem塞进res切片里
// 	for _, val := range nums {
// 		tem += val
// 		res = append(res, tem)
// 	}

// 	//返回带有所有sum的切片
// 	return res
// }

func runningSum(nums []int) []int {
	// 循环数组，给后一个数赋值，值等于本身加上前一个数字
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	// 返回改造后的数组
	return nums
}
