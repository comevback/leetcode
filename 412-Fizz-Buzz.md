# 412. Fizz Buzz

Easy

Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
 

Example 1:
> Input: n = 3
Output: ["1","2","Fizz"]

Example 2:
> Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]

Example 3:
> Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
 

Constraints:

> 1 <= n <= 104

## My First Try (√)
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	// 定义一个切片，容量为n
	var arr []string = make([]string, 0, n)
	// 遍历n次，每次都把n的FizzBuzz情况输入到arr里
	for i := 0; i < n; i++ {
		arr = append(arr, CheckFB(i+1))
	}

	// 返回得到的切片
	return arr
}

// 定义检查一个整数是否是FizzBuzz的函数
func CheckFB(num int) string {
	// 如果既不被3整除又不被5整除，返回改数的字符串形式
	if num%3 != 0 && num%5 != 0 {
		return strconv.Itoa(num)
		// 如果既被3整除又被5整除，返回"FizzBuzz"
	} else if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	} else {
		// 如果只被3整除，返回"Fizz"
		if num%3 == 0 {
			return "Fizz"
			// 剩下只有一种情况：如果只被5整除，返回"Buzz"
		} else {
			return "Buzz"
		}
	}
}
```