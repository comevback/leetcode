# 88. Merge Sorted Array

Easy

Hint
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

 

Example 1:
> Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:
> Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3:
> Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
 

Constraints:

> nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109
 

Follow up: Can you come up with an algorithm that runs in O(m + n) time?

--- 

## My First Try (X)

```go
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
```

## Answer
```go
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
```