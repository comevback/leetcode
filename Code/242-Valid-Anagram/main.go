package main

func main() {

}

// 1.异或
// func isAnagram(s string, t string) bool {
// 	var XOR rune
// 	for _, v := range s + t {
// 		XOR ^= v
// 	}

// 	if XOR == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// 如果遇到 s = "aa"， t = "bb"这种情况，就会判断错误

// 2.map
func isAnagram(s string, t string) bool {
	hsMap := make(map[rune]int)

	for _, v := range s {
		hsMap[v] += 1
	}

	for _, v := range t {
		hsMap[v] -= 1
	}

	for _, val := range hsMap {
		if val != 0 {
			return false
		}
	}

	return true
}

// 3.ASCII值相差
// func isAnagram1(s string, t string) bool {
// 	sum := 0
// 	for _, v := range s {
// 		sum += int(v)
// 	}

// 	for _, v := range t {
// 		sum -= int(v)
// 	}

// 	if sum == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// 如果遇到 s = "ac"， t = "bb"这种情况，就会判断错误

// 4. Array代替map
func isAnagram1(s string, t string) bool {
	chars := make([]int, 26)

	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
