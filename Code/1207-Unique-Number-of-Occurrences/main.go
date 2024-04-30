package main

import "fmt"

func main() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	res := uniqueOccurrences(arr)
	fmt.Println(res)
}

// 1.哈希map O(n)-O(n)
// 统计数组中每个元素出现的次数,记录到map中
// 再统计map中每个元素出现的次数,记录到map中，如果有重复的次数，返回false
func uniqueOccurrences(arr []int) bool {
	var hsMap map[int]int = make(map[int]int)
	var boolMap map[int]bool = make(map[int]bool)

	for _, v := range arr {
		hsMap[v] += 1
	}

	for _, value := range hsMap {
		if boolMap[value] {
			return false
		}
		boolMap[value] = true
	}

	return true
}
