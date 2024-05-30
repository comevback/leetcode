package main

import "fmt"

func main() {
	res := countPrimes(100)
	fmt.Println(res)
}

// 找出小于 n 的所有素数的个数，筛选法
func countPrimes(n int) int {
	// 设一个长度为 n 的数组，所有元素初始化为 true
	var IsPrime []bool = make([]bool, n)
	for i := range IsPrime {
		IsPrime[i] = true
	}

	// 从 2 开始遍历，如果是素数，则将其倍数全部标记为 false
	for i := 2; i*i < n; i += 1 {
		if IsPrime[i] {
			// 从 i 的平方开始，因为小于 i 的倍数已经被标记过了
			for j := i * i; j < n; j += i {
				IsPrime[j] = false
			}
		}
	}

	// 从2开始遍历数组，统计素数的个数
	res := []int{}
	for i := 2; i < n; i++ {
		if IsPrime[i] {
			res = append(res, i)
		}
	}

	return len(res)
}
