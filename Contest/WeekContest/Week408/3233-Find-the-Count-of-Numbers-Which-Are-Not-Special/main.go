package main

import (
	"fmt"
	"math"
)

func main() {
	res := nonSpecialCount(5, 7)
	fmt.Println(res)
}

func nonSpecialCount(l int, r int) int {
	nums := r - l + 1

	rSqrt := math.Sqrt(float64(r))

	primes := getPrimeNums(int(rSqrt))

	for _, v := range primes {
		if v*v >= l && v*v <= r {
			nums -= 1
		}
	}

	return nums
}

func getPrimeNums(num int) []int {
	IsPrime := make([]bool, num+1)
	for i := range IsPrime {
		IsPrime[i] = true
	}

	for i := 2; i*i <= num; i++ {
		if IsPrime[i] {
			for j := i * i; j <= num; j += i {
				IsPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= num; i++ {
		if IsPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// func getPrimeNumsWithout(num int) int {
// 	IsPrime := make([]bool, num)
// 	for i := range IsPrime {
// 		IsPrime[i] = true
// 	}

// 	for i := 2; i*i < num; i++ {
// 		if IsPrime[i] {
// 			for j := i * i; j < num; j += i {
// 				IsPrime[j] = false
// 			}
// 		}
// 	}

// 	res := 0
// 	for i := 2; i < num; i++ {
// 		if IsPrime[i] {
// 			res += 1
// 		}
// 	}

// 	return res
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n <= 3 {
// 		return true
// 	}
// 	if n%2 == 0 || n%3 == 0 {
// 		return false
// 	}
// 	for i := 5; i*i <= n; i += 6 {
// 		if n%i == 0 || n%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
