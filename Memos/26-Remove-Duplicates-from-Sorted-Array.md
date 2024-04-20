# 26. Remove Duplicates from Sorted Array

Easy

Hint
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
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
> Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
> Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
 

Constraints:
> 1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.

---

## My First Try (√) 双指针(双下标)

```go
package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	// 设下标为1，当前元素currentNum为nums[0]，因为第一个元素肯定是要保存的，所以修改后的nums[k]从第二个元素开始查询
	k := 1
	var currentNum int = nums[0]

	// 循环nums，如果这个元素和当前元素不等，放入nums[k]中，同时把这个元素变成当前元素，k+1，如果相同则略过，
	for _, val := range nums {
		if currentNum != val {
			currentNum = val
			nums[k] = val
			k += 1
		}
	}

	// 另一种循环方式
	// for i := 0; i < len(nums); i++ {
	// 	if currentNum != nums[i] {
	// 		currentNum = nums[i]
	// 		nums[k] = nums[i]
	// 		k += 1
	// 	}

	// 	fmt.Println(nums)
	// }

	// 返回k的，也就是修改后的切片长度
	return k
}
```

## 因为输出过多，想到另一种方法，如果nums里的值特别多，但是属于重复率极高，实际上没几种，而这几种我们可以知道是哪几种，这时候可以循环所有可能的数，如果一旦数出现在nums里，则添加进nums[k]

```go
// 循环所有可能的数，如果这个数出现在nums里，则添加进nums[k]
func removeDuplicates2(nums []int) int {

	k := 0

	for i := -100; i <= 100; i++ {
		// 用二分查找法
		if binSearch(nums, i) != -1 {
			nums[k] = i
			k += 1
		}
	}

	return k
}

// 二分查找法实现
func binSearch(nums []int, target int) int {

	// 定义一左一右两个下标
	l, r := 0, len(nums)-1

	// 如果左下标还小于右下标，继续循环
	for l <= r {
		// 得到中间下标
		m := (l + (r - l)) / 2
		// 如果中间下标对应的值等于要找的值，直接返回中间下标
		if target == nums[m] {
			return m
			// 如果中间下标对应的值小于要找的值，要找的值可能位于左侧，所以右下标变为中间下标-1
		} else if target < nums[m] {
			r = m - 1
			// 如果中间下标对应的值大于要找的值，要找的值可能位于右侧，所以左下标变为中间下标+1
		} else {
			l = m + 1
		}
	}

	// 如果循环结束还找不到，返回-1
	return -1
}
```