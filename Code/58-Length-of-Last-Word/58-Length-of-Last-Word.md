# 58. Length of Last Word（字符串分割）

Easy

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.


Example 1:
> Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

Example 2:
> Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.

Example 3:
> Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.
 

Constraints:
> 1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.

---

# Code
```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord2(str))
}

// 1. ===============================================  用strings.Fields分割多个空格  ===============================================
func lengthOfLastWord(s string) int {
	res := strings.Fields(s)[len(strings.Fields(s))-1]
	return len(res)
}

// 2. ==========================================  用bufio.scanner  预定义的bufio.ScanWords  =========================================
func lengthOfLastWord1(s string) int {
	reader := strings.NewReader(s)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	resLength := len(res[len(res)-1])
	return resLength
}

// 3. ===============================================   用bufio.scanner 自定义的splitter  ==========================================
func lengthOfLastWord2(s string) int {
	reader := strings.NewReader(s)
	scanner := bufio.NewScanner(reader)
	scanner.Split(spaceSplitter)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	resLength := len(res[len(res)-1])
	return resLength
}

func spaceSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for start < len(data) && data[start] == ' ' {
		start++
	}

	// ---------- 执行完这一部分后，就会重新调用spaceSplitter --------------------
	for i := start; i < len(data); i++ {
		if data[i] == ' ' {
			return i + 1, data[start:i], nil
		}
	}
	// ----------------------------------------------------------------------
	
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}
```