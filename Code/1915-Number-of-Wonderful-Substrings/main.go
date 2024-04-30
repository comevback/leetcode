package main

import "fmt"

func main() {
	str := string("fihhicigaaahcbgdabfbch")
	fmt.Println(wonderfulSubstrings1(str))
}

// 1.多重循环，时间复杂度太高
func wonderfulSubstrings1(word string) int64 {
	length := len(word)
	var subStringNum int64 = int64(length)
	subStringLength := 2

	// 每次大循环，寻找subStringLength个长度的子串
	for subStringLength <= length {
		// 在原字符串中遍历，找到每个subStringLength长的子串，再对其进行判定
		for i := 0; i <= length-subStringLength; i++ {
			subString := word[i : i+subStringLength]
			fmt.Println(subString)
			// 如果是wonderful string，subStringNum加一
			res := checkWonderful(subString)
			if res {
				subStringNum += 1
			}
		}
		subStringLength += 1
	}

	return subStringNum
}

// 在子串里，设一个map，如果出现了一次，加入到字典，如果第二次碰见这个字符，消除掉，最后看字典长度是否小于2，如果是，则是wonderful string
func checkWonderful(str string) bool {
	var hsMap map[rune]bool = make(map[rune]bool)
	for _, v := range str {
		if hsMap[v] {
			delete(hsMap, v)
		} else {
			hsMap[v] = true
		}
	}
	if len(hsMap) < 2 {
		return true
	} else {
		return false
	}
}

// =================================================================================================================================
// 2.分治法
// 对于一个长度大于1的string，其wonderful string数量等于自身加两个子串的wonderful string数量
// 所以每个string的数量都用checkWonderful(str) + checkWonderful(str[:length-1]) + checkWonderful(str[1:]) - wonderfulSubstrings(word[1:len(word)-1])
func wonderfulSubstrings2(word string) int64 {
	if len(word) == 1 {
		return 1
	}
	if len(word) == 2 {
		if checkWonderful(word) {
			return 3
		} else {
			return 2
		}
	} else {
		if checkWonderful(word) {
			return 1 + wonderfulSubstrings(word[:len(word)-1]) + wonderfulSubstrings(word[1:]) - wonderfulSubstrings(word[1:len(word)-1])
		} else {
			return wonderfulSubstrings(word[:len(word)-1]) + wonderfulSubstrings(word[1:]) - wonderfulSubstrings(word[1:len(word)-1])
		}
	}
}

// 3.用二进制掩码进行统计
// 不直接检查某个子串是否是wonderful string，而是在遍历字符串中字符的过程中，如果word[:i+5]表达的掩码和word[:i]相同，那么word[i:i+5]肯定是一个wonderful string
// 如果word[:i+5]表达的掩码和word[:j]相差一个字符，那么word[j:i+5]也肯定是一个wonderful string

// 所以我们每遇到一个字符串组成的新的掩码，就把它加入到map中，然后遍历map，如果map中有一个掩码和当前掩码相同，那么就找到了一个wonderful string，然后map[此掩码] + 1
// 如果之前这个掩码出现了n次，那么此时word[:i]和之前的n个元素截成的子串都是wonderful string

func wonderfulSubstrings(word string) int64 {
	// 只有1024种可能的掩码状态（2^10）
	countMap := make(map[int]int64)
	prefixMask := 0
	countMap[0] = 1 // 空前缀掩码出现一次
	var result int64 = 0

	for i := 0; i < len(word); i++ {
		// 更新当前字符对应的位掩码
		prefixMask ^= 1 << (word[i] - 'a')

		// 统计所有字符出现偶数次的子串数量，在i之前，有多少个掩码和当前掩码相同
		result += countMap[prefixMask]

		// 统计恰有一个字符出现奇数次的子串数量，在i之前，有多少个掩码和当前掩码相差一个字符
		for j := 0; j < 10; j++ {
			result += countMap[prefixMask^(1<<j)]
		}

		// 更新当前掩码的出现次数
		countMap[prefixMask]++
	}

	return result
}
