# 344. Reverse String
Solved
Easy
Topics
Companies
Hint
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
> Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
> Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
 

Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.

---

# Code
```go
package main

func main() {

}

// 双指针
func reverseString(s []byte) {
	var head, tail int = 0, len(s) - 1

	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head += 1
		tail -= 1
	}
}
```
