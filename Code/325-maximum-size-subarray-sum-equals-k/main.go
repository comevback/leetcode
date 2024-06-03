package main

import "fmt"

func main() {
	nums := []int{1, -1, 5, -2, 3}
	k := 3
	res := MaxSubArrayLen(nums, k)
	fmt.Println(res)
}

func MaxSubArrayLen(nums []int, k int) int {
	maxLength := 0
	hsMap := make(map[int][]int)
	sum := 0
	hsMap[0] = append(hsMap[0], -1)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		need := sum - k
		if len(hsMap[need]) > 0 {
			length := i - hsMap[need][0]
			if length > maxLength {
				maxLength = length
			}
		}
		hsMap[sum] = append(hsMap[sum], i)
	}

	return maxLength
}
