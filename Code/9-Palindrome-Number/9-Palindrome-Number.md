# 9. Palindrome Number

Solved
Easy
Topics
Companies
Hint
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:

-231 <= x <= 231 - 1

Follow up: Could you solve it without converting the integer to a string?

---

# Code

```go
package main

import "strconv"

func main() {

}

// isPalindrome 使用字符串方法检查一个整数是否为回文
func isPalindrome(x int) bool {
	// 将整数转换为字符串
	numStr := strconv.Itoa(x)
	// 将字符串转换为字符数组
	numArr := []rune(numStr)

	// 如果字符数组的长度小于2，直接返回true
	if len(numArr) < 2 {
		return true
	}

	// 初始化头尾指针
	head, tail := 0, len(numArr)-1

	// 循环比较头尾指针指向的字符
	for head < tail {
		if numArr[head] != numArr[tail] {
			// 如果字符不相等，则不是回文
			return false
		}
		// 移动头尾指针
		head += 1
		tail -= 1
	}

	// 所有字符都相等，则是回文
	return true
}

// isPalindrome1 使用整数反转的方法检查一个整数是否为回文
func isPalindrome1(x int) bool {
	// 负数不可能是回文
	if x < 0 {
		return false
	}

	// 保存原始整数值
	original := x
	// 初始化反转后的整数
	reversed := 0

	// 反转整数
	for x > 0 {
		// 取当前整数的最后一位并加入到反转后的整数中
		reversed = reversed*10 + x%10
		// 去掉当前整数的最后一位
		x = x / 10
	}

	// 比较反转后的整数与原始整数是否相等
	if reversed == original {
		return true
	} else {
		return false
	}
}
```
