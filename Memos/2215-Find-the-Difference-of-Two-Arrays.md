# 2215. Find the Difference of Two Arrays

Easy

Hint
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

 

Example 1:
> Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].

Example 2:
> Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
 

Constraints:
> 1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000

---

```go
package main

import "fmt"

func main() {
	nums1 := []int{3, 2, 1}
	nums2 := []int{2, 6, 4}

	res := findDifference(nums1, nums2)
	fmt.Println(res)
}

// 1.哈希map O(m+n)-O(m+n)
// 统计两个数组的元素种类到map，然后比较两个map的差异
// 把map1的key和map2的key进行比较，如果map1的key不在map2中，说明这个key是nums1中的独有元素，加入到结果列表里，反之亦然
func findDifference(nums1 []int, nums2 []int) [][]int {
	var hsMap1 map[int]bool = make(map[int]bool)
	var hsMap2 map[int]bool = make(map[int]bool)

	var result [][]int = make([][]int, 2)

	for _, v := range nums1 {
		hsMap1[v] = true
	}

	for _, v := range nums2 {
		hsMap2[v] = true
	}

	for key := range hsMap1 {
		if !hsMap2[key] {
			result[0] = append(result[0], key)
		}
	}

	for key := range hsMap2 {
		if !hsMap1[key] {
			result[1] = append(result[1], key)
		}
	}

	return result

}

// =================================================================================================================================
// // 2.排序
// func findDifference1(nums1 []int, nums2 []int) [][]int {
// 	sort(nums1)
// 	sort(nums2)

// 	fmt.Println(nums1)
// 	fmt.Println(nums2)
// 	return [][]int{}
// }

// func sort(nums []int) {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if nums[j] > nums[j+1] {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 			fmt.Println(nums)
// 		}
// 	}
// }
```