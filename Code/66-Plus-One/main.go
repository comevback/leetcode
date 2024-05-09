package main

import "fmt"

func main() {
	digits := []int{8, 9, 9, 9}
	fmt.Println(plusOne2(digits))
}

// 从后往前位运算
func plusOne(digits []int) []int {
	// 如果最后一位不是9，直接加一，其他都不变
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
	} else {
		// 如果最后一位是9，从后往前遍历，如果是9，变为0，如果不是9，加一，如果最后一位是9，变为0，前面加一位1
		add := 1
		// 如果第一位是9，且进位还为1，说明所有位都是9，需要进位，给数组尾部加一个位，所有元素后移一位，然后把数组第一位变为1，退出循环
		for i := len(digits) - 1; i >= 0; i-- {
			if i == 0 && add == 1 && digits[i] == 9 {
				digits[i] = 0
				digits = append(digits, 0)
				for i := 0; i < len(digits)-1; i++ {
					digits[i+1] = digits[i]
				}
				digits[0] = 1
				break
			}
			// 如果是9，变为0，如果不是9，add进位仍然为1，如果加一后不是10，这个位变为加一后的值，add变为0，退出循环
			if digits[i]+add == 10 {
				add = 1
				digits[i] = 0
			} else {
				digits[i] = digits[i] + add
				add = 0
				break
			}
		}
	}
	// 返回数组
	return digits
}

// 2. 翻转数组
func plusOne2(digits []int) []int {
	// 翻转数组
	reverseArray(digits)
	// 进位符，初始化为1，因为一定要在末位加1
	add := 1
	// 从前往后遍历数组
	for i := 0; i < len(digits); i++ {
		// 如果当前位加进位符为10，当前位变为0，进位符变为1
		if digits[i]+add == 10 {
			digits[i] = 0
			add = 1
		} else {
			// 如果当前位加进位符不为10，当前位变为当前位加进位符，进位符变为0，退出循环
			digits[i] += 1
			add = 0
			break
		}
	}

	// 如果最后一位加进位符为10，说明所有位都是9，需要进位，给数组尾部加一个位，所有元素后移一位，然后把数组第一位变为1
	if add == 1 {
		digits = append(digits, 1)
	}

	// 翻转数组
	reverseArray(digits)
	// 返回数组
	return digits
}

// 翻转数组方法
func reverseArray(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
