package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 7, 3, 2, 5}
	res := countSmaller(arr)
	fmt.Println(res)
}

// 1. 我的解法，快速排序，复杂度太高，不行
func countSmaller1(nums []int) []int {
	res := make([]int, len(nums))
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			res[i] = res[i-1]
			continue
		}
		copy(temp[i:], nums[i:])
		res[i] = QuickSort(temp[i:])
	}
	return res
}

func QuickSort(nums []int) int {
	pivot := nums[0]
	index := 0
	compare := len(nums) - 1

	for index < compare {
		if nums[compare] >= pivot { // 相等也要移动，因为右边必须全是小于当前pivot的值
			if compare == index+1 {
				nums[compare], nums[index] = nums[index], nums[compare]
			} else {
				nums[compare], nums[index], nums[index+1] = nums[index+1], nums[compare], nums[index]
			}
			index += 1
		} else {
			compare -= 1
		}
	}
	return len(nums) - index - 1
}

// ***********************************************************************************************************************

// 2.归并排序，在左右两个数组合并的时候，每当左侧数组的元素被取出时，就可以计算这个元素右侧小于它的元素的数量，也就是 rightIndex - (mid + 1)
// 把这个 rightIndex - (mid + 1) 累加到这个元素的 count 中，最终就是这个元素右侧小于它的元素的数量

// Pair 结构体用于封装元素及其在原数组中的索引，当元素在排序过程中移动时，可以通过这个Pair里的Index找到这个元素在Count数组中的位置
// 从而在Count数组的对应位置中更新这个元素右侧小于它的元素的数量
type Pair struct {
	Val   int
	Index int
}

// temp 用于归并排序中暂时存储元素，在一般的归并排序中，用整形数组即可，但这里需要同时存储元素和索引，所以使用Pair结构体
var temp []Pair

// count 用于记录每个元素右侧小于它的元素的数量，返回时直接返回这个数组即可
var count []int

// countSmaller 主函数
func countSmaller(nums []int) []int {
	// 初始化所有类型的数组
	length := len(nums)
	temp = make([]Pair, length)
	count = make([]int, length)
	// 初始化一个Pair类型的数组 arr
	arr := make([]Pair, length)

	// 把原数组 nums 中的每个元素的值和Index， 以Pair形式组合起来，放入 arr 中
	for i, num := range nums {
		arr[i] = Pair{Val: num, Index: i}
	}

	// 使用 arr 数组进行归并排序，排序过程中会更新count数组
	MergeSort(arr, 0, length-1)

	// 返回被更新过的 count 数组
	return count
}

// MergeSort 使用归并排序算法递归地排序数组，同时计算右侧较小元素的数量
func MergeSort(arr []Pair, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	MergeSort(arr, lo, mid)
	MergeSort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

// ** Merge函数 ** 合并两个有序数组
// ** 每当左侧数组的元素被取出时，就可以计算这个元素右侧小于它的元素的数量，也就是 rightIndex - (mid + 1)
// ** 把这个 rightIndex - (mid + 1) 累加到这个元素的 count 中，最终就是这个元素右侧小于它的元素的数量
func merge(arr []Pair, lo int, mid int, hi int) {

	// 把需要排序的数组部分复制到临时数组部分中
	copy(temp[lo:hi+1], arr[lo:hi+1])

	// 定义左右指针，分别指向左半部分的起始位置，右半部分的起始位置
	leftIndex, rightIndex := lo, mid+1
	// 当前指针，指向当前需要填充的位置
	current := lo

	// 当左右两个部分都不越界时，继续比较
	for leftIndex <= mid && rightIndex <= hi {
		// 当要存入的元素是左侧数组的元素时，存入这个元素，并累加这个元素右侧小于它的元素数量
		if temp[leftIndex].Val <= temp[rightIndex].Val {
			arr[current] = temp[leftIndex]
			count[arr[current].Index] += rightIndex - (mid + 1) // ** 当从左数组取元素时，累加这个元素右侧小于它的元素数量
			leftIndex++
		} else {
			arr[current] = temp[rightIndex]
			rightIndex++
		}
		current++
	}

	// ** 处理左侧数组剩余元素，同时累加这个元素右侧小于它的元素数量
	for leftIndex <= mid {
		arr[current] = temp[leftIndex]
		count[arr[current].Index] += rightIndex - (mid + 1) // ** 当从左数组取元素时，累加这个元素右侧小于它的元素数量
		leftIndex++
		current++
	}

	// 处理右侧数组剩余元素
	for rightIndex <= hi {
		arr[current] = temp[rightIndex]
		rightIndex++
		current++
	}
}
