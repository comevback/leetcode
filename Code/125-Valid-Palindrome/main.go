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
