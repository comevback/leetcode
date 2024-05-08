package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr []int = []int{10, 3, 8, 9, 4}
	res := findRelativeRanks(arr)
	fmt.Println(res)
	fmt.Println(len(res))
}

// 1. 排序，map
// 先复制一个数组，进行排序，然后用map存储每一个值的排名 key为值，value为排名
// 遍历原数组，根据map的值，判断排名，然后赋值给结果数组，结果数组是string类型，如果遇到1，2，3则赋值为对应的奖牌字符串
func findRelativeRanks(score []int) []string {

	// 复制数组
	var copyScore []int = make([]int, len(score))
	copy(copyScore, score)
	// 定义结果数组
	res := make([]string, len(score))

	// 插入排序，也可以使用别的排序算法
	// 例如：sort.Ints(copyScore)
	insertSorting(copyScore)
	// 定义map，key为值，value为排名
	var sortMap map[int]int = make(map[int]int, len(score))

	// 遍历排序后的数组，存储排名
	for i, v := range copyScore {
		sortMap[v] = len(score) - i
	}

	// 遍历原数组，根据map的值，判断排名，然后赋值给结果数组
	for i, v := range score {
		value := strconv.Itoa(sortMap[v])
		switch value {
		case "1":
			value = "Gold Medal"
		case "2":
			value = "Silver Medal"
		case "3":
			value = "Bronze Medal"
		}

		res[i] = value
	}

	// 返回结果数组
	return res
}

// 插入排序
func insertSorting(nums []int) {
	// 从第二个元素开始遍历，每个i都是未排名的，i之前的元素都是已排名的
	for i := 1; i < len(nums); i++ {
		// 从i开始，向前遍历，如果i小于i-1，则交换位置
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// ========================================================================================================================
