package main

import "fmt"

func main() {
	var arr []int = []int{3, 3}
	fmt.Println(twoSum2(arr, 6))
}

// 第一种方法，两层遍历
func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		// 第二层，记得从第一层的下一个元素开始，
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	// 如果没有结果，返回一个空的切片
	return []int{}
}

// 第二种方法，用哈希表map
func twoSum2(nums []int, target int) []int {

	// 定义一个map，key的值是元素值，value的值是切片的索引
	var hsMap map[int]int = make(map[int]int)

	// 遍历所有元素，如果在当前循环中，找到了以val的互补数(target-val)为key的map元素，也就是这个值在map中存在，那么返回这两个元素的下标
	for i, val := range nums {

		// 找当前值的互补值，是否在map中，如果在，取其下标和当前遍历的下标i组合，返回
		if j, exist := hsMap[target-val]; exist {
			return []int{i, j}
		}
		//如果没有，把当前值val作为key，i作为value，储存到map中
		hsMap[val] = i
	}

	// 如果没有结果，返回空切片
	return []int{}
}
