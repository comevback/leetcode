package main

import (
	"fmt"
)

func main() {
	nums := 3749
	res := intToRoman(nums)
	fmt.Println(res)
}

// intToRoman converts an integer to a Roman numeral.
func intToRoman(num int) string {
	res := ""
	thousand := num / 1000 // 计算千位上的数字
	num -= 1000 * thousand // 减去千位的部分，更新num
	hundred := num / 100   // 计算百位上的数字
	num -= 100 * hundred   // 减去百位的部分，更新num
	ten := num / 10        // 计算十位上的数字
	num -= 10 * ten        // 减去十位的部分，更新num
	one := num             // 个位上的数字

	// 添加千位的罗马数字
	for i := 0; i < thousand; i++ {
		res += "M"
	}

	// 根据百位的数字添加相应的罗马数字
	if hundred == 9 {
		res += "CM"
	} else if hundred == 4 {
		res += "CD"
	} else {
		if hundred >= 5 {
			res += "D"
			hundred -= 5
		}
		for i := 0; i < hundred; i++ {
			res += "C"
		}
	}

	// 根据十位的数字添加相应的罗马数字
	if ten == 9 {
		res += "XC"
	} else if ten == 4 {
		res += "XL"
	} else {
		if ten >= 5 {
			res += "L"
			ten -= 5
		}
		for i := 0; i < ten; i++ {
			res += "X"
		}
	}

	// 根据个位的数字添加相应的罗马数字
	if one == 9 {
		res += "IX"
	} else if one == 4 {
		res += "IV"
	} else {
		if one >= 5 {
			res += "V"
			one -= 5
		}
		for i := 0; i < one; i++ {
			res += "I"
		}
	}

	return res
}

// 2. solution from leetcode
func intToRoman1(num int) string {
	var res string
	for num >= 1000 {
		res += "M"
		num -= 1000
	}
	if num >= 900 {
		res += "CM"
		num -= 900
	}
	for num >= 500 {
		res += "D"
		num -= 500
	}
	if num >= 400 {
		res += "CD"
		num -= 400
	}
	for num >= 100 {
		res += "C"
		num -= 100
	}
	if num >= 90 {
		res += "XC"
		num -= 90
	}
	for num >= 50 {
		res += "L"
		num -= 50
	}
	if num >= 40 {
		res += "XL"
		num -= 40
	}
	for num >= 10 {
		res += "X"
		num -= 10
	}
	if num >= 9 {
		res += "IX"
		num -= 9
	}
	for num >= 5 {
		res += "V"
		num -= 5
	}
	if num >= 4 {
		res += "IV"
		num -= 4
	}
	for num >= 1 {
		res += "I"
		num -= 1
	}
	return res
}
