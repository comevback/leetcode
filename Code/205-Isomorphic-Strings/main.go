package main

import "fmt"

func main() {
	a := "badc"
	b := "baba"

	res := isIsomorphic2(a, b)
	fmt.Println(res) // 输出 false
}

// 1. isIsomorphic 检查两个字符串是否是同构的
func isIsomorphic(s string, t string) bool {
	typeNum1 := checkType(s)
	typeNum2 := checkType(t)

	// 如果两个字符串中不同字符的数量不相等，则它们不是同构的
	if typeNum1 != typeNum2 {
		return false
	}

	// 用于存储 s 到 t 的字符映射
	hsMap := make(map[byte]byte)

	// 遍历字符串，检查字符映射关系
	for i := 0; i < len(s); i++ {
		if _, exist := hsMap[s[i]]; exist {
			// 如果已经存在映射关系但不匹配，返回 false
			if hsMap[s[i]] != t[i] {
				return false
			}
		} else {
			// 建立新的字符映射
			hsMap[s[i]] = t[i]
		}
	}

	return true
}

// checkType 计算字符串中不同字符的数量
func checkType(str string) int {
	res := 0
	typeMap := make(map[rune]bool)

	for _, v := range str {
		if !typeMap[v] {
			typeMap[v] = true
			res += 1
		}
	}
	return res
}

// 2. isIsomorphic2 检查两个字符串是否是同构的
func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 用于存储 s 到 t 的字符映射
	sToT := make(map[byte]byte)
	// 用于存储 t 到 s 的字符映射
	tToS := make(map[byte]byte)

	// 遍历字符串，检查字符映射关系
	for i := 0; i < len(s); i++ {
		c1, c2 := s[i], t[i]

		if v, ok := sToT[c1]; ok {
			// 如果已经存在映射关系但不匹配，返回 false
			if v != c2 {
				return false
			}
		} else {
			// 建立新的字符映射
			sToT[c1] = c2
		}

		if v, ok := tToS[c2]; ok {
			// 如果已经存在映射关系但不匹配，返回 false
			if v != c1 {
				return false
			}
		} else {
			// 建立新的字符映射
			tToS[c2] = c1
		}
	}

	return true
}
