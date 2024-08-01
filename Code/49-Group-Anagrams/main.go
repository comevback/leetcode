package main

func main() {

}

// groupAnagrams 将数组中的字符串分组，每个组包含彼此是异位词的字符串。
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}                // 用于存储最终的分组结果
	hsMap := make(map[string][]string) // 哈希表，键是排序后的字符串，值是原始字符串数组

	// 遍历输入的每一个字符串
	for _, v := range strs {
		runeArr := []rune(v)                        // 将字符串转换成 rune 切片以便排序
		temp := make([]rune, len(runeArr))          // 创建一个临时切片用于辅助排序
		mergeSort(runeArr, temp, 0, len(runeArr)-1) // 对 rune 切片进行归并排序
		str := string(runeArr)                      // 将排序后的 rune 切片转换回字符串
		hsMap[str] = append(hsMap[str], v)          // 将原字符串添加到对应的哈希表条目
	}

	// 将哈希表中的所有值收集到结果切片中
	for _, val := range hsMap {
		res = append(res, val)
	}

	return res // 返回分组结果
}

// mergeSort 对 rune 切片进行归并排序
func mergeSort(nums []rune, temp []rune, leftIndex int, rightIndex int) {
	if rightIndex <= leftIndex {
		return // 如果当前区间无效或只有一个元素则返回
	}
	mid := leftIndex + (rightIndex-leftIndex)/2     // 计算中点
	mergeSort(nums, temp, leftIndex, mid)           // 递归排序左半部分
	mergeSort(nums, temp, mid+1, rightIndex)        // 递归排序右半部分
	merge(nums, temp, leftIndex, mid+1, rightIndex) // 合并两个已排序的部分
}

// merge 合并两个已排序的部分
func merge(nums []rune, temp []rune, leftIndex int, midIndex int, rightIndex int) {
	copy(temp[leftIndex:rightIndex+1], nums[leftIndex:rightIndex+1]) // 复制当前部分到临时数组
	l, r := leftIndex, midIndex                                      // 初始化两个指针分别指向左右部分的开始
	index := leftIndex                                               // 初始化用于写回结果的指针

	// 遍历左右两部分，将较小的元素写回原数组
	for l < midIndex && r <= rightIndex {
		if temp[l] <= temp[r] {
			nums[index] = temp[l]
			l++
		} else {
			nums[index] = temp[r]
			r++
		}
		index++
	}

	// 将左侧剩余部分写回原数组
	for l < midIndex {
		nums[index] = temp[l]
		l++
		index++
	}

	// 将右侧剩余部分写回原数组
	for r <= rightIndex {
		nums[index] = temp[r]
		r++
		index++
	}
}
