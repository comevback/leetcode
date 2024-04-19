# 1480. Running Sum of 1d Array

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

Example 1:
> Input: nums = [1,2,3,4]
Output: [1,3,6,10]gi
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

Example 2:
> Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Example 3:
> Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]


## My First Try (√)

```go
package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(runningSum(arr1))
}

func runningSum(nums []int) []int {
	// 定义一个整数型切片，长度为0，容量为给出的nums的长度len(nums)
	var res []int = make([]int, 0, len(nums))

	// 定义一个整数，存储目前累加的值
	var tem int = 0

	// 循环nums切片，累加tem，每次的tem塞进res切片里
	for _, val := range nums {
		tem += val
		res = append(res, tem)
	}

	//返回带有所有sum的切片
	return res
}
```

## My Second Try (√)

```go
func runningSum(nums []int) []int {
	// 循环数组，给后一个数赋值，值等于本身加上前一个数字
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	// 返回改造后的数组
	return nums
}
```