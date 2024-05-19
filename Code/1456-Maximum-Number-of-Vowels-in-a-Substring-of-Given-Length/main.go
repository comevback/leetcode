package main

import "fmt"

func main() {
	str := "abciiidef"
	res := maxVowels(str, 3)
	fmt.Println(res)
}

// 1. 暴力遍历 O(M*N)
func maxVowels1(s string, k int) int {
	max := 0
	for i := 0; i < len(s)-k+1; i++ {
		num := CalVowels(s[i : i+k])
		if num > max {
			max = num
		}
	}
	return max
}

// ================================================================================================================
// 2. 贪心算法
func maxVowels(s string, k int) int {
	num := CalVowels(s[:k])
	max := num
	// 这里的循环条件要特别注意，从1开始到len(s)-k-1，因为最后一个值是len(s)-1,子串相隔k-1个元素，所以子串的头是len(s)-k+1
	for i := 1; i < len(s)-k+1; i++ {
		if s[i-1] == 'a' || s[i-1] == 'i' || s[i-1] == 'e' || s[i-1] == 'o' || s[i-1] == 'u' {
			num -= 1
		}
		if s[i+k-1] == 'a' || s[i+k-1] == 'i' || s[i+k-1] == 'e' || s[i+k-1] == 'o' || s[i+k-1] == 'u' {
			num += 1
		}
		if num > max {
			max = num
		}
	}
	return max
}

// 找出一个字符串中的元音数量
func CalVowels(str string) int {
	count := 0
	for _, v := range str {
		if v == 'a' || v == 'i' || v == 'e' || v == 'o' || v == 'u' {
			count += 1
		}
	}
	return count
}
