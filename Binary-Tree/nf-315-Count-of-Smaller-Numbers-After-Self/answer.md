```go
package main

// @lc code=start

// 记录数组的元素值与下标
type Pair struct {
	val int
	id  int
}

// 归并排序使用的临时数组
var temp []Pair

// 记录结果
var count []int

func countSmaller315(nums []int) []int {
	n := len(nums)
	temp = make([]Pair, n)
	count = make([]int, n)

	// 记录元素原始的索引位置，以便在 count 数组中更新结果
	arr := make([]Pair, n)
	for i, val := range nums {
		arr[i] = Pair{val, i}
	}
	sort315(arr, 0, n-1)
	return count
}

// 归并排序
func sort315(arr []Pair, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort315(arr, lo, mid)
	sort315(arr, mid+1, hi)
	merge315(arr, lo, mid, hi)
}

// 合并两个有序数组
func merge315(arr []Pair, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = arr[i]
	}

	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			arr[p] = temp[j]
			j++
		} else if j == hi+1 {
			arr[p] = temp[i]
			i++
			// 更新count数组, 此时右半边的值全部小于当前temp[i]
			// 多次归并时可能会重复计算某一个temp[i], 所以需要累加
			count[arr[p].id] += j - mid - 1
		} else if temp[i].val > temp[j].val {
			arr[p] = temp[j]
			j++
		} else {
			arr[p] = temp[i]
			i++
			// 更新count数组, 此时temp[mid+1]到temp[j]之间的值，都小于temp[i]
			count[arr[p].id] += j - mid - 1
		}
	}
}

// @lc code=end
```