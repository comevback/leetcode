# 66. Plus One

Easy

You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

 

Example 1:
> Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

Example 2:
> Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].

Example 3:
> Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
 

Constraints:
> 1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.

---

# Code
```go
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
```