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
