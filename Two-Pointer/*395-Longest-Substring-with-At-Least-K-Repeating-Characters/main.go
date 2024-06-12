package main

func main() {

}

func longestSubstring(s string, k int) int {
	hsMap := make(map[byte]int)
	windowMap := make(map[byte]int)
	maxlength := 0
	valid := 0

	for i := 0; i < len(s); i++ {
		hsMap[s[i]] += 1
	}

	left, right := 0, 0
	for right < len(s) {
		if hsMap[s[right]] < k {
			windowMap = make(map[byte]int)
			valid = 0
			right += 1
			left = right
			continue
		}

		windowMap[s[right]] += 1
		if windowMap[s[right]] >= k {
			valid += 1
		}

		if valid == len(windowMap) {

		}
	}

	return maxlength
}
