# 13. Roman to Integer

Easy

Hint
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

> Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example 1:
> Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:
> Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:
> Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 

Constraints:
> 1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].

---

# Code
```go
package main

func main() {

}

// 1.我的解法
func romanToInt(s string) int {
	// 定义累加之和
	sum := 0
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 根据每个字符的值进行选择
		switch s[i] {
		case 'M':
			sum += 1000
		case 'D':
			sum += 500
		case 'C':
			// 如果这个字符是C，且后面还有字符，如果后面的字符是M，累加900，i+1
			if len(s[i:]) >= 2 && s[i+1] == 'M' {
				sum += 900
				i += 1
				// 如果后面的字符是D，累加400，i+1
			} else if len(s[i:]) >= 2 && s[i+1] == 'D' {
				sum += 400
				i += 1
			} else {
				sum += 100
			}
		case 'L':
			sum += 50
		case 'X':
			// 如果这个字符是X，且后面还有字符，如果后面的字符是C，累加90，i+1
			if len(s[i:]) >= 2 && s[i+1] == 'C' {
				sum += 90
				i += 1
				// 如果后面的字符是L，累加40，i+1
			} else if len(s[i:]) >= 2 && s[i+1] == 'L' {
				sum += 40
				i += 1
			} else {
				sum += 10
			}
		case 'V':
			sum += 5
		case 'I':
			// 如果这个字符是I，且后面还有字符，如果后面的字符是X，累加9，i+1
			if len(s[i:]) >= 2 && s[i+1] == 'X' {
				sum += 9
				i += 1
				// 如果后面的字符是V，累加4，i+1
			} else if len(s[i:]) >= 2 && s[i+1] == 'V' {
				sum += 4
				i += 1
			} else {
				sum += 1
			}
		}
	}

	// 返回累加之和
	return sum
}

// 2. solution from others in leetcode
func romanToInt2(s string) int {
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// 从后往前遍历，如果前一个字符比后一个字符小，则减去这一个字符的值，否则加上这一个字符的值
	for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}
```