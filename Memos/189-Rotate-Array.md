# 189. Rotate Array

Medium

Hint
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

 

Example 1:
> Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
> Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
 

Constraints:
> 1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
 

Follow up:
> Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?

--- 

```go
package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(arr, 3)
	fmt.Println(arr)
}

// 1.用切片保存中间过程
func rotate(nums []int, k int) {
	// 如果k大于等于length，取余数
	length := len(nums)
	if k >= length {
		k = k % length
	}
	var tempArr1 []int = nums[length-k:]
	var tempArr2 []int = nums[:length-k]
	result := append(tempArr1, tempArr2...)
	copy(nums, result)
}

// 2.转换下标（仍需定义另外的暂存切片）
func rotate1(nums []int, k int) {
	// 如果k大于等于length，取余数
	length := len(nums)
	if k > length {
		k = k % length
	}
	var tempArr = make([]int, length)
	for i := 0; i < length; i++ {
		j := i + k
		if j >= length {
			j = j - length
		}
		tempArr[j] = nums[i]
	}
	copy(nums, tempArr)
}

// =======================================================================================================================
// 3. 翻转数组reverse O(n)-O(1)
func rotate2(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	reverse3(nums, 0, length-1)

	reverse3(nums, 0, k-1)

	reverse3(nums, k, length-1)

}

// 翻转数组reverse O(n)-O(1)
func rotate3(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	reverse2(nums)
	fmt.Println(nums)
	reverse2(nums[0:k])
	fmt.Println(nums)
	reverse2(nums[k+1 : length-1])
	fmt.Println(nums)
}

// 翻转数组功能（可定义范围）
func reverse(arr []int, start int, end int) {
	// 边界检查
	if end <= start {
		return
	}
	length := end - start
	center := (start + end) / 2
	if length > len(arr) {
		panic("invaild")
	}

	// 在给定的范围内,进行翻转，进行对换的两个元素下标加起来等于start+end
	for i := start; i <= center; i++ {
		temp := arr[i]
		arr[i] = arr[(start+end)-i]
		arr[(start+end)-i] = temp
	}
}

// 翻转数组功能，不可定义范围
func reverse2(arr []int) {
	// 边界检查
	length := len(arr)

	for i := 0; i < length/2; i++ {
		temp := arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = temp
	}
}

// 翻转数组功能
func reverse3(arr []int, start int, end int) {
	length := end - start
	if length > len(arr) {
		panic("invaild")
	}
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// =======================================================================================================================

// 4.环状替换
func rotate4(nums []int, k int) {
	n := len(nums)
	k = k % n          // 如果k大于数组长度，则取余
	count := gcd(k, n) // 计算k和n的最大公约数

	// 从start开始，每次移动k步，直到回到start,然后start+1，直到所有元素都移动到正确的位置,共需要count次
	for start := 0; start < count; start++ {
		current := start
		for {
			// 计算下一个位置
			next := (current + k) % n
			// 交换两个位置的元素
			temp := nums[next]
			nums[next] = nums[start]
			nums[start] = temp
			// 移动到下一个位置
			current = next
			// 如果回到了start，结束这一轮
			if start == current {
				break
			}
		}
	}
}

// 计算最大公约数
func gcd(a, b int) int {
	// 计算a和b的最大公约数：辗转相除法，原理是gcd(a,b) = gcd(b,a%b)，直到b为0，
	// 例如gcd(10,5) = gcd(5,0) = 5，或者gcd(10,3) = gcd(3,1) = gcd(1,0) = 1
	for b != 0 {
		// 交换a和b的值
		a, b = b, a%b
	}
	return a
}
```