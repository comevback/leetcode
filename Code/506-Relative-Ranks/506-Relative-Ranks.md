# 506. Relative Ranks（排序）

Easy

You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on. The placement of each athlete determines their rank:

The 1st place athlete's rank is "Gold Medal".
The 2nd place athlete's rank is "Silver Medal".
The 3rd place athlete's rank is "Bronze Medal".
For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").
Return an array answer of size n where answer[i] is the rank of the ith athlete.

 

Example 1:
> Input: score = [5,4,3,2,1]
Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].

Example 2:
> Input: score = [10,3,8,9,4]
Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].

 

Constraints:
> n == score.length
1 <= n <= 104
0 <= score[i] <= 106
All the values in score are unique.

---

# 思路
## 排序后，储存每个值的排名，然后遍历原数组，根据排名输出对应的奖牌或者排名

# 代码
```go
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
```