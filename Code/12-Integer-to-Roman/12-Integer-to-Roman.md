# 12. Integer to Roman

Solved
Medium
Topics
Companies
Seven different symbols represent Roman numerals with the following values:

Symbol Value
I 1
V 5
X 10
L 50
C 100
D 500
M 1000
Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
Given an integer, convert it to a Roman numeral.

Example 1:

Input: num = 3749

Output: "MMMDCCXLIX"

Explanation:

3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
700 = DCC as 500 (D) + 100 (C) + 100 (C)
40 = XL as 10 (X) less of 50 (L)
9 = IX as 1 (I) less of 10 (X)
Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
Example 2:

Input: num = 58

Output: "LVIII"

Explanation:

50 = L
8 = VIII
Example 3:

Input: num = 1994

Output: "MCMXCIV"

Explanation:

1000 = M
900 = CM
90 = XC
4 = IV

Constraints:

1 <= num <= 3999

---

# Code

```go
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
```
