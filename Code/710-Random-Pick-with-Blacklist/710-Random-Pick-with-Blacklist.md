# 710. Random Pick with Blacklist
Solved
Hard
Topics
Companies
You are given an integer n and an array of unique integers blacklist. Design an algorithm to pick a random integer in the range [0, n - 1] that is not in blacklist. Any integer that is in the mentioned range and not in blacklist should be equally likely to be returned.

Optimize your algorithm such that it minimizes the number of calls to the built-in random function of your language.

Implement the Solution class:

Solution(int n, int[] blacklist) Initializes the object with the integer n and the blacklisted integers blacklist.
int pick() Returns a random integer in the range [0, n - 1] and not in blacklist.
 

Example 1:

Input
["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
[[7, [2, 3, 5]], [], [], [], [], [], [], []]
Output
[null, 0, 4, 1, 6, 1, 0, 4]

Explanation
Solution solution = new Solution(7, [2, 3, 5]);
solution.pick(); // return 0, any integer from [0,1,4,6] should be ok. Note that for every call of pick,
                 // 0, 1, 4, and 6 must be equally likely to be returned (i.e., with probability 1/4).
solution.pick(); // return 4
solution.pick(); // return 1
solution.pick(); // return 6
solution.pick(); // return 1
solution.pick(); // return 0
solution.pick(); // return 4
 

Constraints:

1 <= n <= 109
0 <= blacklist.length <= min(105, n - 1)
0 <= blacklist[i] < n
All the values of blacklist are unique.
At most 2 * 104 calls will be made to pick.

---

# Code
```go
package main

import (
	"math/rand"
)

func main() {

}

// type Solution struct {
// 	blackMap map[int]bool
// 	arr      []int
// 	useMap   bool
// 	max      int
// }

// func Constructor(n int, blacklist []int) Solution {
// 	hsMap := make(map[int]bool)
// 	arr := []int{}
// 	var useMap bool
// 	for _, v := range blacklist {
// 		hsMap[v] = true
// 	}

// 	if len(blacklist) > n/2 {
// 		useMap = false
// 		for i := 0; i < n; i++ {
// 			if !hsMap[i] {
// 				arr = append(arr, i)
// 			}
// 		}
// 	} else {
// 		useMap = true
// 	}

// 	return Solution{
// 		blackMap: hsMap,
// 		arr:      arr,
// 		useMap:   useMap,
// 		max:      n,
// 	}
// }

// func (this *Solution) Pick() int {
// 	if this.useMap {
// 		index := rand.Intn(this.max)
// 		for this.blackMap[index] {
// 			index = rand.Intn(this.max)
// 		}
// 		return index
// 	} else {
// 		index := rand.Intn(len(this.arr))
// 		return this.arr[index]
// 	}
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

// type Solution1 struct {
// 	arr []int
// }

// func Constructor1(n int, blacklist []int) Solution1 {
// 	sort.Ints(blacklist)
// 	arr := []int{}
// 	current := 0

// 	for i := 0; i < n; i++ {
// 		if current >= len(blacklist) || i != blacklist[current] {
// 			arr = append(arr, i)
// 		} else {
// 			current += 1
// 		}
// 	}

// 	return Solution1{
// 		arr: arr,
// 	}
// }

// func (this *Solution1) Pick1() int {
// 	index := rand.Intn(len(this.arr))
// 	return this.arr[index]
// }

type Solution struct {
	size         int         // 搜索的范围
	blacklistMap map[int]int // 搜索范围中黑名单的值映射到非搜索范围中白名单的值
}

func Constructor(n int, blacklist []int) Solution {
	// 搜索范围 [0, size)
	size := n - len(blacklist)
	// 记录黑名单的位置
	// 这里的value可以使用struct{}类型，而非布尔类型，这样能更加节约空间
	blackBool := make(map[int]struct{}, len(blacklist))
	for _, b := range blacklist {
		blackBool[b] = struct{}{}
	}
	w := size
	blacklistMap := make(map[int]int, len(blacklist))
	for _, b := range blacklist {
		// 如果 b 为搜索范围中的值，则需要在非搜索范围中找到非黑名单值的映射
		if b < size {
			for w < n {
				// 找到非黑名单的值
				if _, ok := blackBool[w]; ok {
					w++
				} else {
					// 将黑名单的值映射到非黑名单的值
					blacklistMap[b] = w
					w++
					break
				}
			}
		}
	}
	return Solution{
		size:         size,
		blacklistMap: blacklistMap,
	}
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.size)
	// 搜索到的值是黑名单的值，则直接返回所引射的值
	if _, ok := this.blacklistMap[index]; ok {
		return this.blacklistMap[index]
	}
	// 搜索到的值是非黑名单的值，可以直接返回
	return index
}
```