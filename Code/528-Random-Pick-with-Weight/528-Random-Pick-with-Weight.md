# 528. Random Pick with Weight
Solved
Medium
Topics
Companies
You are given a 0-indexed array of positive integers w where w[i] describes the weight of the ith index.

You need to implement the function pickIndex(), which randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it. The probability of picking an index i is w[i] / sum(w).

For example, if w = [1, 3], the probability of picking index 0 is 1 / (1 + 3) = 0.25 (i.e., 25%), and the probability of picking index 1 is 3 / (1 + 3) = 0.75 (i.e., 75%).
 

Example 1:

Input
["Solution","pickIndex"]
[[[1]],[]]
Output
[null,0]

Explanation
Solution solution = new Solution([1]);
solution.pickIndex(); // return 0. The only option is to return 0 since there is only one element in w.
Example 2:

Input
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
Output
[null,1,1,1,1,0]

Explanation
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // return 1. It is returning the second element (index = 1) that has a probability of 3/4.
solution.pickIndex(); // return 1
solution.pickIndex(); // return 1
solution.pickIndex(); // return 1
solution.pickIndex(); // return 0. It is returning the first element (index = 0) that has a probability of 1/4.

Since this is a randomization problem, multiple answers are allowed.
All of the following outputs can be considered correct:
[null,1,1,1,1,0]
[null,1,1,1,1,1]
[null,1,1,1,0,0]
[null,1,1,1,0,1]
[null,1,0,1,0,0]
......
and so on.
 

Constraints:

1 <= w.length <= 104
1 <= w[i] <= 105
pickIndex will be called at most 104 times.

---

# Code
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 3, 4, 5, 6}
	obj := Constructor(nums)
	res := obj.PickIndex()
	fmt.Println(res)
}

// Solution 结构体包含一个整数切片 Arr，用来存储权重的累积和。
type Solution struct {
	Arr []int
}

// Constructor 是 Solution 结构体的构造函数，它接受一个权重数组 w，并返回一个 Solution 实例。
// 这个函数计算权重的累积和并存储在 Arr 中。
func Constructor(w []int) Solution {
	var weights []int = make([]int, len(w))
	sum := 0
	for i, v := range w {
		sum += v
		weights[i] = sum
	}

	return Solution{
		Arr: weights,
	}
}

// PickIndex 是 Solution 类的方法，用于根据权重数组随机选择一个索引。
// 它首先生成一个 [0, Arr 最后一个元素) 范围内的随机数。
// 然后使用二分查找确定随机数落在哪个权重区间。
func (this *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器
	r := rand.Intn(this.Arr[len(this.Arr)-1])

	left, right := 0, len(this.Arr)

	for left < right {
		mid := left + (right-left)/2
		if this.Arr[mid] == r {
			return mid + 1
		} else if this.Arr[mid] < r {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := PickIndex();
 */
```