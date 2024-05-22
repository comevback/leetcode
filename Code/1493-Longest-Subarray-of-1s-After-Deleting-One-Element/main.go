package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	res := longestSubarray(arr)
	fmt.Println(res)
}

// 1.我的解法：维护1元素index队列
// 定义一个容量为2的队列，每当遇到0时，就把0的下一个塞进队列，如果已经有两个元素，把头元素退出
// length的长度：如果队列有两个，说明中间是间隔了一个0的，也就是删除了一个0，长度需要减掉这个0，所以是 length = i - oneIndex[0]
// 否则 length = i - oneIndex[0] + 1
func longestSubarray(nums []int) int {
	var length int
	var maxLength int
	var oneIndex [2]int // 记录了1元素下标的队列

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if len(oneIndex) == 0 {
				oneIndex[0] = i + 1
			} else if len(oneIndex) == 1 {
				oneIndex[1] = i + 1
			} else {
				oneIndex[0] = oneIndex[1]
				oneIndex[1] = i + 1
			}
		}

		if len(oneIndex) == 2 {
			length = i - oneIndex[0]
		} else {
			length = i - oneIndex[0] + 1
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
