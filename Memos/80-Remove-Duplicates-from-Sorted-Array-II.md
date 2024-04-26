# 80. Remove Duplicates from Sorted Array II

Medium

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

Example 1:
> Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
> Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.

It does not matter what you leave beyond the returned k (hence they are underscores).
 

Constraints:
>1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


## My First Try (√) 从后往前遍历

```go
package main

import "fmt"

func main() {
	var arr = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(arr)
}

// 从后往前遍历，用copy方法修改原切片
func removeDuplicates(nums []int) int {
	// 设两个下标，i指向倒数第三个元素，j指向最后一个元素
	i, j := len(nums)-3, len(nums)-1
	// k记录返回的切片下标，也就是去重后的长度
	k := len(nums)

	// 从后往前遍历
	for i >= 0 {
		// 如果i和j指向的元素相等，就去掉j指向的元素
		if nums[i] == nums[j] {
			copy(nums[j:], nums[j+1:])
			k -= 1
		}
		i -= 1
		j -= 1
		fmt.Println(nums)
	}
	return k
}
```

## 从前往后遍历 直接操作

```go
// 从前往后遍历 直接操作
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}
	k := 2
	for i := 2; i < length; i++ {
		//注意！这里是比较 nums[i] 和 nums[k-2]。
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k += 1
		}
		fmt.Println(nums)
	}
	return k
}
```

## Using map 统计出现次数

```go
// 用map记录元素出现次数
func removeDuplicates(nums []int) int {
	// 定义一个map，key是切片nums中某一个元素的值，value是这个元素出现的次数
	m := make(map[int]int)
	// cout记录去重后的长度
	cout := 0

	// 遍历nums，统计每个元素出现的次数，如果次数小于等于2，就将这个元素放到nums的前面，然后cout加1
	for i := range nums {
		m[nums[i]]++
		if m[nums[i]] < 3 {
			nums[cout] = nums[i]
			cout++
		}
	}

	return cout
}
```


