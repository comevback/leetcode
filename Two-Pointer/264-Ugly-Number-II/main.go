package main

import (
	"fmt"
	"math"
)

func main() {
	res := nthUglyNumber(11)
	fmt.Println(res)
}

// 1.我的方法，计算出一定范围内所有丑数，然后查找，复杂度太高，不行
func nthUglyNumber1(n int) int {
	res := []int{}
	used := make([]bool, 1691)

	dividePrimesTimes(1690, &used)

	for i := 1; i <= 1690; i++ {
		if !used[i] {
			res = append(res, i)
			if len(res) == n {
				break
			}
		}
	}

	return res[len(res)-1]
}

func dividePrimesTimes(max int, used *[]bool) {
	primes := findPrimes(max)
	primes = primes[3:]

	for _, v := range primes {
		for j := v; j <= max; j += v {
			(*used)[j] = true
		}
	}
}

func findPrimes(max int) []int {
	IsPrime := make([]bool, max+1)
	for i := range IsPrime {
		IsPrime[i] = true
	}
	res := []int{}

	for i := 2; i*i <= max; i++ {
		if IsPrime[i] {
			for j := i * i; j <= max; j += i {
				IsPrime[j] = false
			}
		}
	}

	for i := 2; i <= max; i++ {
		if IsPrime[i] {
			res = append(res, i)
		}
	}
	return res
}

// ********************************************************************************************************************
// 动态规划
func nthUglyNumber(n int) int {
	// 创建一个1-index的数组，从第二项开始储存丑数
	ugly := make([]int, n+1)
	current := 1 // 丑数数组索引从1开始

	// 创建三个值，分别代表三个能被2，3，5整除的整数队列，初始值都为1
	value1, value2, value3 := 1, 1, 1
	// 三个队列的索引都从1开始
	ptr1, ptr2, ptr3 := 1, 1, 1

	// 循环
	for current <= n {
		// 每次取三队头的最小值，放入丑数数组
		min := getMin(value1, value2, value3)
		ugly[current] = min
		current += 1

		// 和最小值相等的，都乘以相应倍数
		// ! 关键部分： 不能用else if，因为如果当前最小值等于多个队列的最小值，所有相等的值都要乘相应倍数，以免重复写入
		if min == value1 {
			value1 = 2 * ugly[ptr1]
			ptr1 += 1
		}
		if min == value2 {
			value2 = 3 * ugly[ptr2]
			ptr2 += 1
		}
		if min == value3 {
			value3 = 5 * ugly[ptr3]
			ptr3 += 1
		}
	}

	// 返回最后一项
	return ugly[n]

}

// 取三个数中最小值
func getMin(n1 int, n2 int, n3 int) int {
	var min int = math.MaxInt

	if n1 < min {
		min = n1
	}

	if n2 < min {
		min = n2
	}

	if n3 < min {
		min = n3
	}

	return min
}
