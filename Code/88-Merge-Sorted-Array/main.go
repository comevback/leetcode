package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}

	merge(arr1, 3, arr2, 3)
	//fmt.Println(arr1)
}

// 错误原因：没有修改原本nums1的值，而是新拷贝了切片
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]

	// 定义一个容量为m+n的切片
	var newArr []int = make([]int, m+n, m+n)

	// 循环m+n次，每次找nums1和nums2中第一个元素，取更小那一个，填进newArr里
	for i := 0; i < m+n; i++ {
		var minNum int
		if len(nums1) == 0 && len(nums2) == 0 {
			break
		}
		if len(nums1) == 0 {
			minNum = nums2[0]
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			minNum = nums1[0]
			nums1 = nums1[1:]
		} else {
			if nums1[0] < nums2[0] {
				minNum = nums1[0]
				nums1 = nums1[1:]
			} else if nums1[0] >= nums2[0] {
				minNum = nums2[0]
				nums2 = nums2[1:]
			}
		}
		newArr[i] = minNum
		fmt.Println(newArr)
	}

	// 把nums1替换成newArr
	nums1 = newArr
}

// answer,从后往前合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1的下标设为i
	i := m - 1
	// nums2的下标设为j
	j := n - 1
	// nums1（合并后）的下标设为k
	k := m + n - 1

	// nums还有元素时，就继续循环
	for j >= 0 {

		// 如果这一轮，nums1还有值，并且大于nums2，就把这个nums1的值放在nums的k位，反之则把nums2的值放k位，放了哪个arr的值，对应下标就减一，最后k也减一，寻找下一个k位的数据。
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i -= 1
		} else {
			nums1[k] = nums2[j]
			j -= 1
		}
		k -= 1
	}
}
