# 27. Remove Element

Easy

Hint
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

 

Example 1:
> Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
> Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
 

Constraints:
> 0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100

---

## My First Try (√) 复制切片，遍历，填入原切片

```go
package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2))
}

// 设一个新切片，满足条件的填入新切片
func removeElement(nums []int, val int) int {
	// 原本的切片长度
	oriLen := len(nums)
	// 要移除的数出现次数
	var count int
	// 复制原切片
	var temp []int = make([]int, len(nums))
	copy(temp, nums)
	// 把原切片清零
	nums = nums[0:0]

	// 在复制的切片里循环，如果遇到要移除的数，不管，计数加一，如果不是要移除的数，加入到原切片里
	for idx, v := range temp {
		if v == val {
			count += 1
		} else {
			nums = append(nums, temp[idx])
		}
		fmt.Println(nums)
	}

	// 返回移除后到切片长度
	return oriLen - count
}
```

## My Second Try (X) 用哈希表表示频率，没有移除原切片中的值
```go
// 用哈希表map来存储出现频率
func removeElement2(nums []int, val int) int {

	var hsMap map[int]int = make(map[int]int)

	for i, v := range nums {
		if v == val {
			hsMap[val] += 1
			nums[i] = 0
		} else {
			hsMap[v] += 1
		}
		fmt.Println(nums)
	}

	return len(nums) - hsMap[val]
}
```

## Answer 双下标 只有符合条件（不等于要移除的值）的元素，才能从头放进数组里，等于筛选了一遍
```go
// 双下标
func removeElement3(nums []int, val int) int {
	// 修改后的下标
	k := 0
	// 循环，碰到不是要移除的值，就放进nums[k]里，遇到是，则不管
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
		fmt.Println(nums)
	}
	// 返回修改后的下标k，也就是原切片中不是val的值的个数
	return k
}
```