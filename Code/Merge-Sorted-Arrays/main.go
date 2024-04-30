package main

import "fmt"

func main() {
	arr1 := []int{0, 3, 4, 31}
	arr2 := []int{4, 6, 30}
	fmt.Println(mergeSortedArrays1(arr1, arr2))
}

// 1.双指针 O(n+m)-O(n+m)
func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	length1 := len(nums1)
	length2 := len(nums2)

	// 如果其中一个数组为空，直接返回另一个数组
	if length1 == 0 {
		return nums2
	}
	if length2 == 0 {
		return nums1
	}

	m, n := 0, 0

	// 预先分配内存
	var result []int = make([]int, 0, length1+length2)

	// nums1和nums2中的指针都没有到达末尾时，取出两个数组中最小的值，放入result中，然后指针后移
	for length1 > m && length2 > n {
		if nums1[m] < nums2[n] {
			result = append(result, nums1[m])
			m += 1
		} else {
			result = append(result, nums2[n])
			n += 1
		}
	}

	// 如果nums1中还有剩余元素，直接放入result中，如果nums2中还有剩余元素，直接放入result中
	if length1 > m {
		result = append(result, nums1[m:]...)
	} else {
		result = append(result, nums2[n:]...)
	}

	return result
}

// 2.取出两组中最小值删除，直到其中一个数组为空 O(n+m)-O(n+m)
func mergeSortedArrays1(nums1 []int, nums2 []int) []int {
	length1 := len(nums1)
	length2 := len(nums2)

	// 如果其中一个数组为空，直接返回另一个数组
	if length1 == 0 {
		return nums2
	}
	if length2 == 0 {
		return nums1
	}

	// 用于存放结果
	var res []int = make([]int, 0, length1+length2)

	// 取出两个数组中最小的值，放入res中，然后删除，直到其中一个数组为空
	for len(nums1) != 0 && len(nums2) != 0 {
		if nums1[0] < nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}

	// 如果nums1中还有剩余元素，直接放入res中，如果nums2中还有剩余元素，直接放入res中
	if len(nums1) > 0 {
		res = append(res, nums1...)
	} else {
		res = append(res, nums2...)
	}

	return res
}

// 3.递归 O(n+m)-O(n+m)
func mergeSortedArrays2(nums1 []int, nums2 []int) []int {

	// 如果其中一个数组为空，直接返回另一个数组
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}

	// 递归取出两个数组中最小的值，放入res中，然后递归调用
	if nums1[0] < nums2[0] {
		return append([]int{nums1[0]}, mergeSortedArrays2(nums1[1:], nums2)...)
	} else {
		return append([]int{nums2[0]}, mergeSortedArrays2(nums1, nums2[1:])...)
	}
}
