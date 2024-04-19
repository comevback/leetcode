package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {

	// 定义ransomNote和magazine的map，key类型时byte，因为string用切片分解出的每个字符，类型为byte
	var dicNote map[byte]int = map[byte]int{}
	var dicMaga map[byte]int = map[byte]int{}

	// 遍历ransomNote和magazine，把每个字符都往各自的map里装一次，每次遇到相同字符，该key下的value就加一
	// 得到在ransomNote和magazine中，每个字符的出现次数
	for i := 0; i < len(ransomNote); i++ {
		dicNote[ransomNote[i]] += 1
	}

	for i := 0; i < len(magazine); i++ {
		dicMaga[magazine[i]] += 1
	}

	// 遍历ransomNote的map，如果ransomNote中的所有key在magazine中都有，而且值大于等于在ransomNote中的值，则返回true，否则返回false
	for key, value := range dicNote {
		if dicMaga[key] >= value {
		} else {
			return false
		}
	}
	return true
}
