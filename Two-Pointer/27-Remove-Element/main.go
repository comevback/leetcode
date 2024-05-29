package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement3(arr, 2))
}

// 设一个新切片，满足条件的填入新切片，不是原地修改
func removeElement1(nums []int, val int) int {
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
	}
	// 返回修改后的下标k，也就是原切片中不是val的值的个数
	return k
}

// *******************************************************************************************************************
// 双指针，原地修改
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow += 1
		}
		fast += 1
	}

	return slow
}
