package main

import "fmt"

func main() {
	var arr = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(arr)
	fmt.Println(arr)
}

// 用copy方法修改原切片 O(n^2) 不好
func removeDuplicates(nums []int) int {
	// 设两个下标，i指向倒数第三个元素，j指向最后一个元素
	i, j := len(nums)-3, len(nums)-1
	// k记录去重后的长度
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
		//fmt.Println(nums)
	}
	return k
}

// 从前往后遍历 直接操作
func removeDuplicates1(nums []int) int {
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

// 用map记录元素出现次数
func removeDuplicates2(nums []int) int {
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
