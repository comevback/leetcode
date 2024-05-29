package main

import "fmt"

func main() {
	str := "ccc"
	res := longestPalindrome(str)
	fmt.Println(res)
}

// 寻找字符串中最长的回文子串
func longestPalindrome1(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		str := palindromic1(s, i)
		if len(str) > len(res) {
			res = str
		}
	}
	return res
}

// 从给定的索引开始扩展，找到最长的回文子串
func palindromic1(str string, index int) string {
	left, right := index, index
	res := ""

	// 扩展回文子串，考虑回文子串长度为奇数的情况和偶数的情况
	// 默认为奇数，但是在扩展的过程中，如果发现左右字符不相等，考虑是否可以切换为偶数情况，如果满足条件，切换为偶数
	for left >= 0 || right < len(str) {
		if left >= 0 && right < len(str) && str[left] == str[right] {
			res = str[left : right+1]
			left -= 1
			right += 1
		} else if left >= 0 && str[left] == str[right-1] {
			if str[left:index] == reverse(str[index:right]) {
				res = str[left:right]
				left -= 1
			} else {
				break
			}
		} else if right < len(str) && str[left+1] == str[right] {
			if str[left+1:index+1] == reverse(str[index+1:right+1]) {
				res = str[left+1 : right+1]
				right += 1
			} else {
				break
			}
		} else {
			break
		}
	}

	return res
}

// 反转字符串
func reverse(str string) string {
	strArr := []byte(str)
	head, tail := 0, len(str)-1

	for head < tail {
		strArr[head], strArr[tail] = strArr[tail], strArr[head]
		head += 1
		tail -= 1
	}

	res := string(strArr)
	return res
}

// ********************************************************************************************************************
// 把判断回文串中心是一个字符还是两个字符的情况提到主函数中，减少逻辑复杂度

// 寻找字符串中最长的回文子串
func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		// 考虑回文子串长度为奇数的情况
		str1 := palindromic(s, i, i)
		// 考虑回文子串长度为偶数的情况
		str2 := palindromic(s, i, i+1)
		if len(str1) > len(res) {
			res = str1
		}
		if len(str2) > len(res) {
			res = str2
		}
	}
	return res
}

// 从给定的左右索引开始扩展，找到最长的回文子串
func palindromic(str string, left int, right int) string {
	res := ""

	// 扩展回文子串
	for left >= 0 && right < len(str) {
		if str[left] == str[right] {
			res = str[left : right+1]
			left -= 1
			right += 1
		} else {
			break
		}
	}

	return res
}
