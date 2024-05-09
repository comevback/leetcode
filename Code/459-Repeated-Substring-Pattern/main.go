package main

import "fmt"

func main() {
	var str string = "abac"
	fmt.Println(repeatedSubstringPattern1(str))
}

// ==================================================================================================================
// 1.异或运算
// func repeatedSubstringPattern(s string) bool {
// 	var XOR rune

// 	for _, v := range s {
// 		XOR ^= v
// 	}

// 	return XOR == 0
// }
// s = "ababab"的情况下失败

// ==================================================================================================================
// 2.检测每个字符出现次数，如果有不同的则返回false，如果有，按照这个次数作为window滑动，看是否是重复子串
// func repeatedSubstringPattern(s string) bool {

// 	if len(s) == 0 || len(s) == 1 {
// 		return false
// 	}

// 	hsMap := make(map[rune]int)

// 	for _, v := range s {
// 		hsMap[v] += 1
// 	}

// 	var times int = hsMap[rune(s[0])] // Convert s[0] to rune
// 	for _, val := range hsMap {
// 		if val != times {
// 			return false
// 		}
// 	}

// 	var str string = s[:len(hsMap)]
// 	for i := 0; i < len(s)-len(hsMap)+1; i += len(hsMap) {
// 		if s[i:i+len(hsMap)] != str {
// 			return false
// 		}
// 	}

// 	return true
// }

// s = "abaababaab"的情况失败

// ==============================================  构造双倍数组，去头尾，查找元素组  ====================================================
func repeatedSubstringPattern(s string) bool {
	str := s + s
	str = str[1 : len(str)-1]
	for i := 0; i < len(str)-len(s)+1; i++ {
		if str[i:i+len(s)] == s {
			return true
		}
	}
	return false
}

// ==========================================  暴力遍历每个长度的子串，时间复杂度太高，错误  ==============================================
func repeatedSubstringPattern1(s string) bool {
	for length := 1; length <= len(s)/2; length++ {
		sum := ""
		for i := 0; i < len(s)/length; i++ {
			sum += s[:length]
		}
		if sum == s {
			return true
		}
	}
	return false
}
