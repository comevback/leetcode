package main

import "fmt"

func main() {
	s1 := "enz"
	s2 := "emhqwznesno"
	res := checkInclusion(s1, s2)
	fmt.Println(res)
}

// 1. 我的解法
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	vaild := 0

	for i := 0; i < len(s1); i++ {
		need[s1[i]] += 1
	}

	left, right := 0, 0

	// 滑动窗口
	for right < len(s2) {
		// 如果右指针指向的字符在need中，也就是在s1中，那么就把这个字符加入到window中
		if need[s2[right]] > 0 {
			// 如果window中这个字符的数量等于need中这个字符的数量，那么vaild就加1，表示这个字符已经满足了，同时右指针右移
			window[s2[right]] += 1
			if window[s2[right]] == need[s2[right]] {
				vaild += 1
			}
			right += 1

			// 如果左右指针之间的距离大于s1的长度，那么就要移动左指针，如果左指针指向的字符在need中，那么就要把这个字符从window中移除
			if right-left > len(s1) {
				if need[s2[left]] > 0 {
					// 如果window中这个字符的数量等于need中这个字符的数量，那么vaild就减1，表示这个字符不满足了
					if window[s2[left]] == need[s2[left]] {
						vaild -= 1
					}
					window[s2[left]] -= 1
				}
				left += 1
			}

			// 如果vaild等于need的长度，那么就表示s2中包含s1的排列
			if vaild == len(need) {
				return true
			}
		} else {
			// 如果右指针指向的字符不在need中，那么就把左右指针都移动到右指针的下一个位置，重新开始
			vaild = 0
			window = make(map[byte]int)
			right += 1
			left = right
		}
	}

	return false
}

// other solution from leetcode
func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := [26]int{}
	s2Count := [26]int{}

	// Populate the counts for s1 and the first window of s2
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	// Function to compare if two count arrays are the same
	matches := func() bool {
		for i := 0; i < 26; i++ {
			if s1Count[i] != s2Count[i] {
				return false
			}
		}
		return true
	}

	// Slide the window over s2
	for i := len(s1); i < len(s2); i++ {
		if matches() {
			return true
		}
		s2Count[s2[i]-'a']++
		s2Count[s2[i-len(s1)]-'a']--
	}

	// Check the last window
	return matches()
}
