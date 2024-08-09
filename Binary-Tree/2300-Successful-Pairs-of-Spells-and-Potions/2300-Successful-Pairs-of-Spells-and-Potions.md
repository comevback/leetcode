# 2300. Successful Pairs of Spells and Potions

Solved
Medium
Topics
Companies
Hint
You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.

Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

Example 1:

Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
Explanation:

- 0th spell: 5 \* [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
- 1st spell: 1 \* [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
- 2nd spell: 3 \* [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
  Thus, [4,0,3] is returned.
  Example 2:

Input: spells = [3,1,2], potions = [8,5,8], success = 16
Output: [2,0,2]
Explanation:

- 0th spell: 3 \* [8,5,8] = [24,15,24]. 2 pairs are successful.
- 1st spell: 1 \* [8,5,8] = [8,5,8]. 0 pairs are successful.
- 2nd spell: 2 \* [8,5,8] = [16,10,16]. 2 pairs are successful.
  Thus, [2,0,2] is returned.

Constraints:

n == spells.length
m == potions.length
1 <= n, m <= 105
1 <= spells[i], potions[i] <= 105
1 <= success <= 1010

---

# Code

```go
package main

import "fmt"

func main() {
	spells := []int{5, 1, 3}                   // 定义一个数组，存储各个魔法的强度
	potions := []int{1, 2, 3, 4, 5}            // 定义一个数组，存储各个药水的强度
	res := successfulPairs(spells, potions, 7) // 调用函数计算成功的配对
	fmt.Println(res)                           // 输出结果
}

// successfulPairs 函数计算可以成功配对的魔法和药水数量
func successfulPairs(spells []int, potions []int, success int64) []int {
	temp := make([]int, len(potions))           // 创建一个临时数组用于排序
	mergeSort(potions, temp, 0, len(potions)-1) // 对药水进行归并排序

	res := make([]int, len(spells)) // 创建结果数组
	for i, v := range spells {
		num := getPairNums(potions, v, success) // 计算每种魔法可以成功配对的药水数量
		res[i] = num
	}

	return res
}

// getPairNums 函数返回可以与指定魔法成功配对的药水数量
func getPairNums(potions []int, spell int, success int64) int {
	if spell*potions[len(potions)-1] < int(success) {
		return 0
	} else if spell*potions[0] >= int(success) {
		return len(potions)
	}

	left, right := 0, len(potions)-1 // 初始化二分查找的边界

	for right >= left {
		mid := left + (right-left)/2 // 计算中间索引
		if potions[mid]*spell >= int(success) {
			right = mid - 1
		} else if potions[mid]*spell < int(success) {
			left = mid + 1
		}
	}

	return len(potions) - left
}

// mergeSort 函数实现归并排序，用于排序药水数组
func mergeSort(nums []int, temp []int, left int, right int) {
	if right <= left {
		return
	}

	mid := left + (right-left)/2          // 计算中间索引
	mergeSort(nums, temp, left, mid)      // 递归排序左半部分
	mergeSort(nums, temp, mid+1, right)   // 递归排序右半部分
	merge(nums, temp, left, mid+1, right) // 合并两个有序的子数组
}

// merge 函数合并两个有序的子数组
func merge(nums []int, temp []int, left int, mid int, right int) {
	copy(temp[left:right+1], nums[left:right+1]) // 拷贝原数组的元素到临时数组
	l, r := left, mid
	current := left

	for l < mid && r <= right {
		if temp[l] <= temp[r] {
			nums[current] = temp[l]
			l++
		} else {
			nums[current] = temp[r]
			r++
		}
		current++
	}

	for l < mid {
		nums[current] = temp[l]
		l++
		current++
	}

	for r <= right {
		nums[current] = temp[r]
		r++
		current++
	}
}
```
