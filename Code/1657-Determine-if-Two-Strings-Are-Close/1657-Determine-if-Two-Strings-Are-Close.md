# 1657. Determine if Two Strings Are Close

Medium

Hint
Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

Example 1:
> Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"

Example 2:
> Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

Example 3:
> Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"
 

Constraints:
> 1 <= word1.length, word2.length <= 105
word1 and word2 contain only lowercase English letters.

---

# Code
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	word1 := "abc"
	word2 := "bca"

	fmt.Println(closeStrings(word1, word2))
}

// 1. 比较两个字符串中，出现过的字符种类，和出现过的次数
func closeStrings(word1 string, word2 string) bool {
	// 创建两个哈希表，用于存储两个字符串中字符出现次数
	hsMap1 := make(map[rune]int)
	hsMap2 := make(map[rune]int)

	// 如果两个字符串长度不相等，直接返回false
	if len(word1) != len(word2) {
		return false
	}

	// 遍历两个字符串，将字符出现次数存入哈希表
	for _, v := range word1 {
		if _, exist := hsMap1[v]; !exist {
			hsMap1[v] = 1
		} else {
			hsMap1[v] += 1
		}
	}

	for _, v := range word2 {
		if _, exist := hsMap2[v]; !exist {
			hsMap2[v] = 1
		} else {
			hsMap2[v] += 1
		}
	}

	// 创建一个哈希表，用于存储哈希表中每个value出现的次数
	var valueMap map[int]int = make(map[int]int)

	// 遍历第一个哈希表，如果第二个哈希表中不存在该key，直接返回false，同时将value出现次数存入valueMap
	for key, value := range hsMap1 {
		if _, exist := hsMap2[key]; !exist {
			return false
		}
		valueMap[value] += 1
	}

	// 遍历第二个哈希表，每个value在valueMap中减一，如果有某个值减到0以下，直接返回false
	for _, value := range hsMap2 {
		if valueMap[value] < 1 {
			return false
		} else {
			valueMap[value] -= 1
		}
	}

	// 如果上述条件都满足，返回true
	return true
}

// =====================================================================================================================
// 2. 数组和排序
// 基本思想还是对比两个字符串出现的字符种类，和出现的次数
func closeStrings1(word1 string, word2 string) bool {
	// 将两个字符串转换成字节切片
	wordOneBytes := []byte(word1)
	wordTwoBytes := []byte(word2)

	// 创建两个大小为128的整数数组来存储每个字符的出现频率
	// ASCII码总共有128个字符，这样可以直接使用字符的ASCII值作为索引
	setOne := make([]int, 128)
	setTwo := make([]int, 128)

	// 遍历第一个字符串的字节切片，统计每个字符的出现次数
	for i := 0; i < len(wordOneBytes); i++ {
		setOne[wordOneBytes[i]]++
	}

	// 遍历第二个字符串的字节切片，统计每个字符的出现次数
	for i := 0; i < len(wordTwoBytes); i++ {
		setTwo[wordTwoBytes[i]]++
	}

	// 检查两个字符串中的字符是否完全相同
	for i := 0; i < 128; i++ {
		// 如果字符在一个字符串中出现而在另一个字符串中没有出现，则直接返回false
		if setOne[i] > 0 && setTwo[i] == 0 {
			return false
		}

		// 如果字符在一个字符串中出现而在另一个字符串中没有出现，则直接返回false
		if setTwo[i] > 0 && setOne[i] == 0 {
			return false
		}
	}

	// 对字符出现的频率进行排序
	sort.Ints(setOne)
	sort.Ints(setTwo)

	// 比较排序后的频率，如果有不同，则这两个字符串不是“close strings”
	for i := 0; i < 128; i++ {
		if setOne[i] != setTwo[i] {
			return false
		}
	}

	// 所有的检查都通过了，返回true
	return true
}
```