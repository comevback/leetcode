package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes3(nums)
	fmt.Println(nums)
}

// ================================================================================================================================

// 1.暴力双循环 O(n^2)-O(1)
func moveZeroes1(nums []int) {
	length := len(nums)

	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if nums[i] == 0 {
				nums = append(nums[:i], nums[i+1:]...)
				nums = append(nums, 0)
			}
		}
	}
}

// ================================================================================================================================

// 2.从后往前循环，冒泡 O(n*k)-O(1)
// 假设数组的长度为 n，并且有 k 个0分布在数组的前半部分，那么每次操作的平均复杂度为 O(n)，因为可能需要移动近 n 个元素。由于需要对每个0执行这样的操作，总的时间复杂度将接近于 O(n*k)。
func moveZeroes(nums []int) {
	length := len(nums)

	// 从后往前循环，如果是0，就把数组后面的元素都往前移动一位，然后把0塞到最后
	for i := length - 1; i >= 0; i-- {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
}

// ================================================================================================================================

// 3.双指针 O(n)-O(1)
// 从前往后遍历，把每个不为零的数都往前面堆
func moveZeroes2(nums []int) {
	length := len(nums)
	head := 0

	// 从前往后遍历
	for i := 0; i < length; i++ {
		// 如果不为0，就把这个数和head指向的数交换，然后head+1
		if nums[i] != 0 {
			nums[i], nums[head] = nums[head], nums[i]
			head += 1
		}
	}
}

// 4.双指针优化 O(n)-O(1)
// 从前往后遍历，把每个不为零的数都往前面堆，但不是交换，而是直接赋值
func moveZeroes3(nums []int) {
	length := len(nums)
	head := 0

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[head] = nums[i]

			if i != head {
				nums[i] = 0
			}

			head += 1
		}
	}
}