package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 7, 9, 0, 0, 0}
	arr2 := []int{2, 5, 6}

	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1)
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
	// nums1 = newArr
	copy(nums1[0:], newArr[0:])
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 定义一个新切片来储存nums1中的值，如果不这样，直接在nums1中操作，过程中会覆盖掉还没操作的数据
	var temp = make([]int, m)
	copy(temp, nums1[:m])
	// 定义下标，三个切片均从0开始
	i, j, k := 0, 0, 0

	// 如果temp和nums2都没取到底，每次都把temp和nums2中的第i或j个元素拿出来比较，小的放进nums1里，然后i或j加一
	for j < n && i < m {
		// 如果i小于m并且temp中的值小于nums2中的值，就把temp中的值放进nums1里，反之把nums2中的值放进nums1里
		if temp[i] <= nums2[j] {
			nums1[k] = temp[i]
			i += 1
		} else {
			nums1[k] = nums2[j]
			j += 1
		}
		k += 1
	}

	// 如果nums2中还有数据，就把nums2中的数据放进nums1里
	for j < n {
		nums1[k] = nums2[j]
		j += 1
		k += 1
	}

	// 如果temp中还有数据，就把temp中的数据放进nums1里
	for i < m {
		nums1[k] = temp[i]
		i += 1
		k += 1
	}
}

// answer,从后往前合并
func merge3(nums1 []int, m int, nums2 []int, n int) {
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

// =====================================================================================================================
// 3. 不依赖外部空间的解法
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:m+n], nums2)
	quickSort(nums1)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])

	return mergeIndex(left, right)
}

func mergeIndex(nums1 []int, nums2 []int) []int {
	res := []int{}
	p1, p2 := 0, 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] <= nums2[p2] {
			res = append(res, nums1[p1])
			p1 += 1
		} else {
			res = append(res, nums2[p2])
			p2 += 1
		}
	}

	if p1 < len(nums1) {
		res = append(res, nums1[p1:]...)
	}

	if p2 < len(nums2) {
		res = append(res, nums2[p2:]...)
	}

	return res
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	key := len(nums) - 1
	head := 0
	for head < key {
		if nums[head] > nums[key] {
			if head+1 == key {
				nums[head], nums[key] = nums[key], nums[head]
			} else {
				nums[head], nums[key], nums[key-1] = nums[key-1], nums[head], nums[key]
			}
			key -= 1
		} else {
			head += 1
		}
	}

	quickSort(nums[:key])
	if key < len(nums)-1 {
		quickSort(nums[key+1:])
	}
}

// =====================================================================================================================
