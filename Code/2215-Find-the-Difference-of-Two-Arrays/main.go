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
