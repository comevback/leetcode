# 125. Valid Palindrome

Easy

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.


Example 1:
> Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
> Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
> Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 

Constraints:
> 1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.


---

# Code
```go
package main

import (
	"fmt"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	res := isPalindrome(str)
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	var arr []byte
	// s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' || s[i] <= 'z' && s[i] >= 'a' {
			arr = append(arr, s[i])
		}
		if s[i] <= 'Z' && s[i] >= 'A' {
			arr = append(arr, s[i]+32)
		}
	}

	head, tail := 0, len(arr)-1

	for head < tail {
		if arr[head] != arr[tail] {
			return false
		}
		head += 1
		tail -= 1
	}

	return true
}
```