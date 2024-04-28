# 1071. Greatest Common Divisor of Strings（数学解法，用最大公约数找最大公约字符串）（贪心算法）

Easy

Hint
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

 

Example 1:
> Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
> Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
> Input: str1 = "LEET", str2 = "CODE"
Output: ""
 

Constraints:
> 1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.

---

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(fingGcd("ABCABC"))
}

// 1.找最大公约字符串 最坏O(n^2 m)-O(n^2)
func gcdOfStrings(str1 string, str2 string) string {
	// 如果两个字符串相等，直接返回str1
	if str1 == str2 {
		return str1
	}
	// 获取两个字符串的长度，取最大值为maxLength，取最小值为minLength
	// minLength再取maxLength的一半和minLength两个字符串长度的最小值，因为如果更长的字符串的一半都比短的字符串长，那么gcd最长也就是短的字符串
	// 如果短字符串的长度比maxLength的一半还长，那么gcd肯定比maxLength的一半还短，所以循环时取这个最短值作为边界
	length1 := len(str1)
	length2 := len(str2)
	maxLength := max(length1, length2)
	minLength := min(length1, length2)
	minLength = min(maxLength/2, minLength)

	// 遍历两个字符串，找到最大公约字符串
	var gcd string = ""
	for i := 0; i < minLength; i++ {
		// 首先两个字符串的前i+1个字符要相等
		if str1[:i+1] == str2[:i+1] {
			// 然后两个字符串的前i+1个字符要能够整除两个字符串的长度
			if length1%(i+1) == 0 && length2%(i+1) == 0 {
				// 判断这两个字符串的前i+1个字符是否能够拼接成两个字符串，是否是约数
				g1 := strings.Repeat((str1[:i+1]), length1/(i+1))
				g2 := strings.Repeat((str2[:i+1]), length2/(i+1))
				// 如果能够拼接成两个字符串，那么这个字符串就是最大公约字符串，更新gcd
				if g1 == str1 && g2 == str2 {
					gcd = str1[:i+1]
				}
			}
		}
	}
	return gcd
}

// 找一个字符串中的最大重复字符串
func fingGcd(str string) string {
	length := len(str)
	var gcd string = ""
	for i := 0; i < length/2; i++ {
		g := strings.Repeat((str[:i+1]), length/(i+1))
		if g == str {
			gcd = string(str[:i+1])
		}
	}
	return gcd
}

// ===============================================================================================================================
// 数学解法
// 1.计算长度的 GCD：首先计算两个字符串长度的最大公约数。
// 2.构造可能的最长重复子字符串：然后，以这个长度为单位截取其中一个字符串的前 GCD 个字符作为候选的最长重复子字符串。
// 3.验证：验证这个候选子字符串重复若干次后，是否能够分别构成原来的两个字符串。
func gcdOfStrings1(str1 string, str2 string) string {
	// 如果两个字符串是有公约的，那么两个字符串拼接起来的字符串一定是两个字符串的公约字符串
	// 换句话说，两个字符串拼接起来，无论是str1+str2还是str2+str1，都是一样的
	// 所以如果str1+str2 != str2+str1，那么两个字符串没有公约，直接返回空字符串
	if str1+str2 != str2+str1 {
		return ""
	}

	length1 := len(str1)
	length2 := len(str2)

	// 计算两个字符串长度的最大公约数
	gcdLength := gcd(length1, length2)

	// ！！！因为str1+str2 == str2+str1，所以两个字符串的最大公约字符串一定是str1或者str2的前gcdLength个字符
	return str1[:gcdLength]
}

// 找两个数的最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
```