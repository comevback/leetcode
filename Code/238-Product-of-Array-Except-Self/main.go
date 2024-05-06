package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{-1, 1, 0, -3, 3}
	res := productExceptSelf(nums)
	fmt.Println(res)
}

// ===============================================================================================================================
// 1. 多次分别循环，从前往后循环，计算每个元素前面的乘积，从后往前循环，计算每个元素后面的乘积，最后都结果数组中，每个元素等于它前面的乘积乘以后面的乘积
func productExceptSelf1(nums []int) []int {
	var zeroMap map[int]bool = make(map[int]bool)
	var zeroPos int
	var signal bool
	var res []int = make([]int, len(nums))
	productFront := 1
	productBack := 1

	for i, v := range nums {
		if v == 0 {
			if zeroMap[v] {
				return make([]int, len(nums))
			} else {
				zeroPos = i
				zeroMap[v] = true
				signal = true
			}
		} else {
			res[i] = productFront
			productFront *= v
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			res[i] *= productBack
			productBack *= nums[i]
		}
	}

	if signal {
		res[zeroPos] = productFront
		return res
	} else {
		return res
	}

}

// ===============================================================================================================================
// 优化版本
// 不使用额外的map，直接使用变量记录0的个数
func productExceptSelf(nums []int) []int {
	// 0的个数
	var zeroCount int = 0
	// 前面的乘积
	var frontProduct int = 1
	// 后面的乘积
	var backProduct int = 1
	// 定义结果数组，长度和输入数组一样，每个元素都初始化为1
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	// 定义头尾指针，头指针从0开始，尾指针从数组长度-1开始
	var head, tail int = 0, len(nums) - 1

	// 当头指针小于数组长度时，循环
	for head < len(nums) {
		// 如果0的个数等于2，直接返回全0数组
		if zeroCount == 2 {
			return make([]int, len(nums))
		}
		// 如果头指针的值等于0，0的个数加1
		if nums[head] == 0 {
			zeroCount += 1
		}
		// 前面的乘积乘以头指针的值，赋值给结果数组的头指针位置
		res[head] *= frontProduct
		frontProduct *= nums[head]
		// 后面的乘积乘以尾指针的值，赋值给结果数组的尾指针位置
		res[tail] *= backProduct
		backProduct *= nums[tail]
		// 头指针加1，尾指针减1
		head += 1
		tail -= 1
	}

	return res
}

// ========================================================================================================================
// review  2024.5.6
func review(nums []int) []int {
	frontProduct, backProduct := 1, 1
	head, tail := 0, len(nums)-1
	zeroCount := 0

	var res []int = make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	for head < len(nums) {
		if nums[head] == 0 || nums[tail] == 0 {
			if zeroCount == 1 {
				return make([]int, len(nums))
			} else {
				zeroCount += 1
			}
		}

		res[head] *= frontProduct  //此时frontProduct还没有乘nums[head]元素本身，所以是前面所有元素之积
		frontProduct *= nums[head] //这时才乘nums[head]元素本身

		res[tail] *= backProduct  //此时backProduct还没有乘nums[tail]元素本身，所以是前面所有元素之积
		backProduct *= nums[tail] //这时才乘nums[tail]元素本身

		head += 1
		tail -= 1
	}

	return res
}
