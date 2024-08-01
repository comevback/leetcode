# 692. Top K Frequent Words

Solved
Medium
Topics
Companies
Given an array of strings words and an integer k, return the k most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.

Example 1:

Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Explanation: "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:

Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.

Constraints:

1 <= words.length <= 500
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
k is in the range [1, The number of unique words[i]]

Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?

---

# Code

```go
package main

import "fmt"

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	// words := []string{"aaa", "aa", "a"}
	res := topKFrequent2(words, 2)
	fmt.Println(res)
}

// ********************************************** 归并排序解法 ***********************************************
// topKFrequent 返回数组中频率最高的k个单词。
func topKFrequent(words []string, k int) []string {
	hsMap := make(map[string]int) // 创建哈希表来存储每个单词及其出现次数

	// 遍历单词列表，统计每个单词的出现频率
	for _, v := range words {
		hsMap[v] += 1
	}

	itemArr := []Item{} // 创建一个Item切片来存储单词和对应的频率

	// 将哈希表中的单词和频率转换为Item结构，并添加到切片中
	for key, value := range hsMap {
		newItem := Item{
			str:   key,
			times: value,
		}
		itemArr = append(itemArr, newItem)
	}

	temp := make([]Item, len(itemArr))          // 创建临时数组用于归并排序
	mergeSort(itemArr, temp, 0, len(itemArr)-1) // 对Item切片进行归并排序

	res := []string{} // 初始化结果切片
	// 从排序后的切片中选择频率最高的k个单词
	for i := len(itemArr) - 1; i >= len(itemArr)-k; i-- {
		res = append(res, itemArr[i].str)
	}

	return res // 返回结果
}

// Item 用于存储单词及其出现频率
type Item struct {
	times int
	str   string
}

// mergeSort 对Item切片进行归并排序
func mergeSort(arr []Item, temp []Item, l int, r int) {
	if l >= r {
		return // 如果区间不合理，则返回
	}

	m := l + (r-l)/2             // 计算中点
	mergeSort(arr, temp, l, m)   // 递归排序左半部分
	mergeSort(arr, temp, m+1, r) // 递归排序右半部分
	merge(arr, temp, l, m+1, r)  // 合并两个有序部分
}

// merge 合并两个已排序的部分
func merge(arr []Item, temp []Item, l int, m int, r int) {
	copy(temp[l:r+1], arr[l:r+1]) // 复制当前区间到临时数组
	l1, l2 := l, m
	index := l

	// 合并两部分，同时考虑频率和字典序
	for l1 < m && l2 <= r {
		if temp[l1].times < temp[l2].times {
			arr[index] = temp[l1]
			l1 += 1
		} else if temp[l1].times > temp[l2].times {
			arr[index] = temp[l2]
			l2 += 1
		} else if compLexico(temp[l1].str, temp[l2].str) == temp[l1].str { // 如果频率相同，比较字典序，字典序小的放后面，因为要取频率最高的k个单词
			arr[index] = temp[l2]
			l2 += 1
		} else {
			arr[index] = temp[l1]
			l1 += 1
		}
		index += 1
	}

	// 复制剩余部分
	for l1 < m {
		arr[index] = temp[l1]
		l1 += 1
		index += 1
	}

	for l2 <= r {
		arr[index] = temp[l2]
		l2 += 1
		index += 1
	}
}

// *********************************************** 桶排序解法 **********************************************
func topKFrequent1(words []string, k int) []string {
	hsMap := make(map[string]int) // 创建哈希表来存储每个单词及其出现次数
	res := []string{}

	// 遍历单词列表，统计每个单词的出现频率
	for _, v := range words {
		hsMap[v] += 1
	}

	// 创建桶，桶的索引表示单词出现的频率，每个桶（元素）都是一个字符串切片，用于存储出现频率相同的单词
	// 总共有len(words) + 1个桶，因为单词出现的频率最大为len(words)，如果分配len(words)个桶，则当所有元素都是同一个单词时，频率是len(words)，会越界
	bucket := make([][]string, len(words)+1)

	// 把哈希表中的单词根据频率放入桶中
	for key, value := range hsMap {
		bucket[value] = append(bucket[value], key)
	}

	// 从桶中取出频率最高的k个单词，如果桶中有多个单词，按字典序排序
	for i := len(bucket) - 1; i >= 0 && len(res) < k; i-- {
		if len(bucket[i]) > 0 {
			temp := make([]string, len(bucket[i]))
			mergeSortLex(bucket[i], temp, 0, len(bucket[i])-1)
			for _, v := range bucket[i] {
				if len(res) < k {
					res = append(res, v)
				} else {
					break
				}
			}
		}
	}

	return res
}

// mergeSort 对Item切片进行归并排序
func mergeSortLex(arr []string, temp []string, l int, r int) {
	if l >= r {
		return // 如果区间不合理，则返回
	}

	m := l + (r-l)/2                // 计算中点
	mergeSortLex(arr, temp, l, m)   // 递归排序左半部分
	mergeSortLex(arr, temp, m+1, r) // 递归排序右半部分
	mergeLex(arr, temp, l, m+1, r)  // 合并两个有序部分
}

// merge 合并两个已排序的部分
func mergeLex(arr []string, temp []string, l int, m int, r int) {
	copy(temp[l:r+1], arr[l:r+1]) // 复制当前区间到临时数组
	l1, l2 := l, m
	index := l

	// 合并两部分，同时考虑频率和字典序
	for l1 < m && l2 <= r {
		if compLexico(temp[l1], temp[l2]) == temp[l1] { // 如果频率相同，比较字典序，字典序小的放后面，因为要取频率最高的k个单词
			arr[index] = temp[l1]
			l1 += 1
		} else {
			arr[index] = temp[l2]
			l2 += 1
		}
		index += 1
	}

	// 复制剩余部分
	for l1 < m {
		arr[index] = temp[l1]
		l1 += 1
		index += 1
	}

	for l2 <= r {
		arr[index] = temp[l2]
		l2 += 1
		index += 1
	}
}

// ******************************************* 优先队列解法 ***************************************************
func topKFrequent2(words []string, k int) []string {
	hsMap := make(map[string]int) // 创建哈希表来存储每个单词及其出现次数
	myPQ := NewPQ()               // 创建优先队列

	// 遍历单词列表，统计每个单词的出现频率
	for _, v := range words {
		hsMap[v] += 1
	}

	// 将哈希表中的单词和频率转换为Item结构，并添加到优先队列中
	for key, value := range hsMap {
		newItem := Item{
			str:   key,
			times: value,
		}
		myPQ.Push(newItem)
	}

	// 从优先队列中取出频率最高的k个单词，也就是适用K次Pop操作
	res := []string{}
	for i := 0; i < k; i++ {
		poped := myPQ.Pop()
		res = append(res, poped.str)
	}

	return res
}

// ItemPQ 优先队列结构
type ItemPQ struct {
	arr []Item
}

// NewPQ 创建一个新的优先队列
func NewPQ() ItemPQ {
	return ItemPQ{
		arr: []Item{},
	}
}

// sink 方法通过下沉调整确保堆的性质，修正删除根节点后可能破坏的堆结构。
func (pq *ItemPQ) sink(index int) {
	for {
		left := index * 2
		right := index*2 + 1
		max := index

		// 如果左子节点的频率大于等于当前节点的频率，或者左子节点的频率等于当前节点的频率且左子节点的字典序小于当前节点的字典序
		// 则选择左子节点作为最大值
		if left < len(pq.arr) && pq.arr[left].times >= pq.arr[max].times {
			if pq.arr[left].times == pq.arr[max].times {
				if compLexico(pq.arr[left].str, pq.arr[max].str) == pq.arr[left].str {
					max = left
				}
			} else {
				max = left
			}
		}

		// 如果右子节点的频率大于等于当前节点的频率，或者右子节点的频率等于当前节点的频率且右子节点的字典序小于当前节点的字典序
		// 则选择右子节点作为最大值
		if right < len(pq.arr) && pq.arr[right].times >= pq.arr[max].times {
			if pq.arr[right].times == pq.arr[max].times {
				if compLexico(pq.arr[right].str, pq.arr[max].str) == pq.arr[right].str {
					max = right
				}
			} else {
				max = right
			}
		}

		// 如果当前节点已经是最大的，停止下沉。
		if max == index {
			break
		}

		// 否则，与最大的子节点交换。
		pq.arr[index], pq.arr[max] = pq.arr[max], pq.arr[index]
		index = max
	}
}

// swim 方法通过上浮调整确保堆的性质，修正插入新节点后可能破坏的堆结构。
func (pq *ItemPQ) swim(index int) {
	for {
		parent := index / 2

		// 如果当前节点的频率大于其父节点的频率，或者当前节点的频率等于其父节点的频率且当前节点的字典序小于其父节点的字典序
		// 则与父节点交换，否则停止上浮
		if parent >= 1 && pq.arr[parent].times < pq.arr[index].times {
			pq.arr[parent], pq.arr[index] = pq.arr[index], pq.arr[parent]
			index = parent
		} else if parent >= 1 && pq.arr[parent].times == pq.arr[index].times {
			if compLexico(pq.arr[parent].str, pq.arr[index].str) == pq.arr[index].str {
				pq.arr[parent], pq.arr[index] = pq.arr[index], pq.arr[parent]
				index = parent
			} else {
				break
			}
		} else {
			break
		}
	}
}

// Push 方法将一个新的值加入到优先队列中，并进行上浮调整。
func (pq *ItemPQ) Push(item Item) {
	if len(pq.arr) < 1 {
		pq.arr = append(pq.arr, Item{})
	}
	pq.arr = append(pq.arr, item)
	pq.swim(len(pq.arr) - 1)
}

// Pop 方法移除并返回优先队列中的最大元素，并进行下沉调整。
func (pq *ItemPQ) Pop() Item {
	if len(pq.arr) < 1 {
		return Item{}
	}

	poped := pq.arr[1]
	pq.arr[1], pq.arr[len(pq.arr)-1] = pq.arr[len(pq.arr)-1], pq.arr[1]
	pq.arr = pq.arr[:len(pq.arr)-1]
	pq.sink(1)
	return poped
}

// ***************************************** 字典序比较函数以及min函数 *****************************************
// compLexico 比较两个字符串的字典序，返回较小的字符串
func compLexico(str1 string, str2 string) string {
	minLen := min(len(str1), len(str2))

	for i := 0; i < minLen; i++ {
		if str1[i] < str2[i] {
			return str1
		} else if str1[i] > str2[i] {
			return str2
		}
	}

	if len(str1) < len(str2) {
		return str1
	} else {
		return str2
	}
}

// min 返回两个整数中较小的一个
func min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	} else {
		return num2
	}
}
```
