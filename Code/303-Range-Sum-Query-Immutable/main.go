package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)
	res := na.SumRange(0, 2)
	fmt.Println(res)
}

type NumArray struct {
	Arr    []int
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefix[i+1] = sum
	}

	return NumArray{
		Arr:    nums,
		prefix: prefix,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
