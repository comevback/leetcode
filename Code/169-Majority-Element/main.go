package main

import (
	"fmt"
)

func main() {
	var arr = []int{2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1, 2}
	fmt.Println(majorityElement1(arr))
}

// 1.用map来保存次数
func majorityElement(nums []int) int {
	marjo := len(nums) / 2
	var hsMap = make(map[int]int)
	for _, v := range nums {
		hsMap[v] += 1
		if hsMap[v] > marjo {
			return v
		}
	}
	return 0
}

// 2.Boyer-Moore 投票算法
func majorityElement1(nums []int) int {
	cand := 0
	count := 0
	// 如果当前count是0，设当前元素为候选人
	for _, v := range nums {
		if count == 0 {
			cand = v
		}
		// 如果当前元素是候选人，票数加一，反之减一
		if v == cand {
			count += 1
		} else {
			count -= 1
		}
	}
	// 最后还能撑到当选的，就过半数了
	return cand
}
