## 二分查找的区间类型

二分查找可以使用两种类型的区间：左闭右开[left, right)和左闭右闭[left, right]。以下是这两种方式的关键规律和区别。

### 左闭右开 [left, right)

- **初始化**:
  - `left = 0`
  - `right = nums.length`（不包含最后一个索引）

- **循环条件**:
  - `left < right`

- **中点更新**:
  - `nums[mid] < target`: `left = mid + 1`
  - `nums[mid] >= target`: `right = mid`

- **结束检查**:
  - 需确认 `left` 是否越界并且 `nums[left] == target`

- **用途**:
  - 适合查找元素的左边界和处理右边界查找，因为右界是开的，方便定位到最后一个满足条件的元素的下一个位置。

### 左闭右闭 [left, right]

- **初始化**:
  - `left = 0`
  - `right = nums.length - 1`（包含最后一个索引）

- **循环条件**:
  - `left <= right`

- **中点更新**:
  - `nums[mid] < target`: `left = mid + 1`
  - `nums[mid] >= target`: `right = mid - 1`

- **结束检查**:
  - 需确认 `right` 是否在合法范围内并且 `nums[right] == target`

- **用途**:
  - 适合查找元素的右边界和精确处理数组的最后一个元素。

# Code
```go
package main

import "fmt"

func main() {
	// 测试数据
	testCases := []struct {
		nums   []int
		target int
		left1  int // 左闭右开左边界
		right1 int // 左闭右开右边界
		left2  int // 左闭右闭左边界
		right2 int // 左闭右闭右边界
	}{
		{[]int{}, 3, -1, -1, -1, -1},                   // 空数组
		{[]int{1}, 1, 0, 0, 0, 0},                      // 单一元素，目标存在
		{[]int{1}, 2, -1, -1, -1, -1},                  // 单一元素，目标不存在
		{[]int{1, 2, 3, 4, 5}, 3, 2, 2, 2, 2},          // 目标元素在中间
		{[]int{1, 2, 2, 2, 3, 4}, 2, 1, 3, 1, 3},       // 目标元素重复多次
		{[]int{1, 2, 3, 3, 3, 4, 5}, 3, 2, 4, 2, 4},    // 目标元素在中间，重复出现
		{[]int{1, 2, 3, 4, 5}, 6, -1, -1, -1, -1},      // 目标元素不存在，大于最大值
		{[]int{2, 2, 2, 2, 2}, 2, 0, 4, 0, 4},          // 所有元素相同，目标存在
		{[]int{2, 2, 2, 2, 2}, 1, -1, -1, -1, -1},      // 所有元素相同，目标不存在
		{[]int{3, 3, 3, 3, 3, 3, 3, 3}, 3, 0, 7, 0, 7}, // 大量重复元素
		{[]int{1, 2, 3, 4, 5}, 0, -1, -1, -1, -1},      // 目标元素不存在，小于最小值
		{[]int{1, 2, 3, 4, 5, 5, 5}, 5, 4, 6, 4, 6},    // 目标元素为最大值，重复出现
	}

	// 执行测试
	for _, tc := range testCases {
		resLeft1 := searchLeftBound1(tc.nums, tc.target)
		resRight1 := searchRightBound1(tc.nums, tc.target)
		resLeft2 := searchLeftBound2(tc.nums, tc.target)
		resRight2 := searchRightBound2(tc.nums, tc.target)

		fmt.Printf("Testing nums: %v, target: %d\n", tc.nums, tc.target)
		fmt.Printf("Expected (Left1, Right1, Left2, Right2): (%d, %d, %d, %d)\n", tc.left1, tc.right1, tc.left2, tc.right2)
		fmt.Printf("Result   (Left1, Right1, Left2, Right2): (%d, %d, %d, %d)\n\n", resLeft1, resRight1, resLeft2, resRight2)
	}
}

// **********************************************  左闭右开  ****************************************************

// 左闭右开区间搜索左边界
// 查找目标值target的最左侧索引
func searchLeftBound1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 空数组直接返回-1
	}
	left, right := 0, len(nums) // 初始化左右指针

	for left < right { // 循环直到左右指针相遇
		mid := left + (right-left)/2 // 计算中间点
		if nums[mid] < target {
			left = mid + 1 // 目标值在右侧，移动左指针
		} else {
			right = mid // 目标值在左侧或等于中值，移动右指针
		}
	}

	if left < len(nums) && nums[left] == target {
		return left // 检查左指针是否有效并返回左边界
	}
	return -1 // 目标值不存在
}

// 左闭右开区间搜索右边界
// 查找目标值target的最右侧索引的后一位
func searchRightBound1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 空数组直接返回-1
	}
	left, right := 0, len(nums) // 初始化左右指针

	for left < right { // 循环直到左右指针相遇
		mid := left + (right-left)/2 // 计算中间点
		if nums[mid] <= target {
			left = mid + 1 // 目标值在右侧，包括中值，移动左指针
		} else {
			right = mid // 目标值在左侧，移动右指针
		}
	}

	if right == 0 || nums[right-1] != target {
		return -1 // 检查右边界是否有效
	}
	return right - 1 // 返回最右边界的前一个位置
}

// **********************************************  左闭右闭  ****************************************************

// 左闭右闭区间搜索左边界
// 查找目标值target的最左侧索引
func searchLeftBound2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 空数组直接返回-1
	}
	left, right := 0, len(nums)-1 // 初始化左右指针

	for left <= right { // 循环直到左右指针重叠
		mid := left + (right-left)/2 // 计算中间点
		if nums[mid] < target {
			left = mid + 1 // 目标值在右侧，移动左指针
		} else {
			right = mid - 1 // 目标值在左侧或等于中值，移动右指针
		}
	}

	if left < len(nums) && nums[left] == target {
		return left // 检查左指针是否有效并返回左边界
	}
	return -1 // 目标值不存在
}

// 左闭右闭区间搜索右边界
// 查找目标值target的最右侧索引
func searchRightBound2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 空数组直接返回-1
	}
	left, right := 0, len(nums)-1 // 初始化左右指针

	for left <= right { // 循环直到左右指针重叠
		mid := left + (right-left)/2 // 计算中间点
		if nums[mid] <= target {
			left = mid + 1 // 目标值在右侧，包括中值，移动左指针
		} else {
			right = mid - 1 // 目标值在左侧，移动右指针
		}
	}

	if right >= 0 && nums[right] == target {
		return right // 检查右指针是否有效并返回右边界
	}
	return -1 // 目标值不存在
}
```