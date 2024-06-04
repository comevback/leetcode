package main

import "fmt"

func main() {
	res := getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}})
	fmt.Println(res)
}

func getModifiedArray(length int, updates [][]int) []int {
	// 初始化一个长度为length的数组arr，初始值为0
	arr := make([]int, length)
	// 创建一个差分数组对象
	diffArr := NewDiff(arr)

	// 遍历每个update，并应用到差分数组中
	for _, v := range updates {
		diffArr.Increase(v[0], v[1], v[2])
	}

	// 从差分数组计算最终的结果数组
	res := diffArr.getNums()
	return res
}

// Diff结构体用于表示差分数组
type Diff struct {
	diff []int
}

// NewDiff创建一个新的差分数组，并初始化
func NewDiff(nums []int) *Diff {
	// 初始化一个与nums长度相同的差分数组
	diff := make([]int, len(nums))
	// 填充差分数组
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	return &Diff{
		diff: diff,
	}
}

// Increase在[start, end]区间内增加value值
func (d *Diff) Increase(start int, end int, value int) {
	// 在start位置增加value
	d.diff[start] += value
	// 在end+1位置减少value，以标记区间结束
	if end+1 < len(d.diff) {
		d.diff[end+1] -= value
	}
}

// getNums从差分数组还原原数组
func (d *Diff) getNums() []int {
	// 初始化结果数组
	res := make([]int, len(d.diff))
	// 第一个元素等于差分数组的第一个元素
	res[0] = d.diff[0]
	// 通过累加差分数组的值计算原数组的每个元素
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}

	return res
}
