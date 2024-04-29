# Longest Word（字符串分割）

Have the function LongestWord(sen) take the sen parameter being passed and return the longest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty. Words may also contain numbers, for example "Hello world123 567"

examples:
```
Input: "fun&!! time"
Output: time
Input: "I love dogs"
Output: love
```

---

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	sen := string("Have the function !LongestWsegegwgwgord(sen) take the sen parameter being passed and return the longest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty. Words may also contain numbers, for example : Hello world123 567")
	fmt.Println(LongestWord1(sen))
}

// ================================================================================================================================
// 1.正则表达式
// 把标点符号和特殊字符替换为空格，然后用空格分割字符串，找出最长的单词
func LongestWord(sen string) string {
	// 用正则表达式匹配所有标点符号和特殊字符
	reg := regexp.MustCompile(`[\p{P}\p{S}]+`) // 匹配所有标点符号和特殊字符
	replaced := reg.ReplaceAllString(sen, " ") // 替换为空格

	// 用一个空格分割字符串，如果有多个空格，会被分割成空字符串，如果让空字符串不被分割，可以用strings.FieldsFunc或strings.Fields
	senArr := strings.Split(replaced, " ")

	longest := 0
	for _, v := range senArr {
		if len(v) > longest {
			longest = len(v)
			sen = v
		}
	}

	return sen
}

// 2.不用正则表达式的方法，用strings.FieldsFunc分割字符串
func LongestWord2(sen string) (maxWord string) {
	// FieldsFunc返回字符串s中满足f(c)的子字符串的列表
	words := strings.FieldsFunc(sen, func(c rune) bool {
		return c < '0' || (c > '9' && c < 'A') || (c > 'Z' && c < 'a') || c > 'z'
	})
	// 遍历words，找出最长的单词
	for _, word := range words {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	return
}

// ================================================================================================================================
// 3.不用正则表达式的方法 （错误！）
// 先用空格分割字符串，然后定义一个length函数，把分割的每个单词传入length函数，返回单词的长度，找出最长的单词
func LongestWord3(str string) string {
	var best string

	words := strings.Split(str, " ")
	for _, w := range words {
		if length(w) > len(best) {
			best = w
		}
	}
	return best
}

// length函数，返回单词的长度，但如果一个单词开头或中间有标点符号，就会被忽略，所以此方法错误
func length(str string) int {
	var size int
	for _, c := range str {
		if unicode.IsLetter(c) {
			size++
		}
	}
	return size
}

// 不用正则表达式的方法（错误！）
func LongestWord1(sen string) string {
	// 用空格分割字符串
	words := strings.Split(sen, " ")
	// maxWord用于存储最长的单词，类型为[]rune
	var maxWord []rune

	// 遍历words
	for _, word := range words {
		// runes用于存储去掉标点符号的单词
		runes := []rune{}
		// 遍历单词的每个字符
		for _, r := range word {
			// 如果字符不是标点符号，就加入runes，但这里如果一个单词开头或中间有标点符号，就会被忽略，所以此方法错误
			if !unicode.IsPunct(r) {
				runes = append(runes, r)
			} else {
				break
			}
		}
		if len(runes) > len(maxWord) {
			maxWord = runes
		}
	}
	return string(maxWord)

}
```