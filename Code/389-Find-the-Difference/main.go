package main

func main() {

}

// =====================================================  map 统计次数  ============================================================
// 1. map遍历统计每个字符的次数
func findTheDifference(s string, t string) byte {
	hsMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hsMap[s[i]] += 1
	}

	for i := 0; i < len(t); i++ {
		hsMap[t[i]] -= 1
	}

	var res byte

	for key, val := range hsMap {
		if val != 0 {
			res = key
		}
	}

	return res
}

// ======================================================   求ASCII值和之差   =======================================================
// 2. 字符转ASCII值，求和求差
// 把两个字符串每个字符的ASCII值加起来，得到和
// 把两个和相减，差值就是多的那个字符的ACSII值，转换回字符就可得到相差的字符
func findTheDifference1(s string, t string) byte {
	var sumS, sumT int

	for _, v := range s {
		sumS += int(v) //每个v是rune类型，这里把rune转化为ascii值，后面吧acsii值转化为byte，这个题的情况可以通用
	}

	for _, v := range t {
		sumT += int(v)
	}

	res := sumT - sumS
	resByte := byte(res)

	return resByte
}

// ======================================================   异或运算   =======================================================
func findTheDifference2(s string, t string) byte {
	var out = t[len(t)-1]

	for i := 0; i < len(s); i++ {
		out ^= s[i] ^ t[i]
	}
	return out
}

// 或

func findTheDifference3(s string, t string) byte {
	var count int
	for _, val := range s + t { // s + t 可以拼接成一个字符串
		count ^= int(val) // 异或每一个字符，因为每两个相同的字符异或为0，所以其他字符都会为0，最后剩下的就是结果字符
	}
	return byte(count)
}
