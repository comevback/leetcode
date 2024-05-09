package main

import (
	"fmt"
	"sort"
)

func main() {
	var happiness []int = []int{1, 2, 3}
	res := maximumHappinessSum(happiness, 2)
	fmt.Println(res)
}

// =====================================================  排序后直接操作  ==========================================================
func maximumHappinessSum1(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var sum int64 = 0
	for i := k; i > 0; i-- {
		if len(happiness) != 0 {
			sum += int64(happiness[len(happiness)-1])
			happiness = happiness[:len(happiness)-1]
			for i := 0; i < len(happiness); i++ {
				if happiness[i] > 0 {
					happiness[i] -= 1
				}
			}
		}
	}

	return sum
}

// =====================================================    ==========================================================
func maximumHappinessSum(happiness []int, k int) int64 {

}
