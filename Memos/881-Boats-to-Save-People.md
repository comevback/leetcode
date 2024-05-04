# 881. Boats to Save People（双指针）（数组代替map/排序）

Medium

You are given an array people where people[i] is the weight of the ith person, and an infinite number of boats where each boat can carry a maximum weight of limit. Each boat carries at most two people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.


Example 1:
> Input: people = [1,2], limit = 3
Output: 1
Explanation: 1 boat (1, 2)

Example 2:
> Input: people = [3,2,2,1], limit = 3
Output: 3
Explanation: 3 boats (1, 2), (2) and (3)

Example 3:
> Input: people = [3,5,3,4], limit = 5
Output: 4
Explanation: 4 boats (3), (3), (4), (5)
 

Constraints:
> 1 <= people.length <= 5 * 104
1 <= people[i] <= limit <= 3 * 104

---

## 解题思路
找出最少的船数，使得每个人都能被载到。每艘船最多载两个人，且载重不超过limit。
所以每艘船要么装一个人，要么装两个人，为了让船数最少，也就是让装的效率最大化，每艘船最好坐满重量负载。
所以思路是寻找两个人的重量之后最好等于limit，如果不等于，差值也要尽可能小。
所以找一重一轻，如果和小于limit，那么两人一起坐船，如果和大于limit，那么重的人单独坐船。

## 代码
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []int{2, 4}
	limit := 5
	fmt.Println(numRescueBoats2(people, limit))
}

// 1.map O(n^2)-O(n) 时间复杂度太高，不可取
// 第一次循环用map记录每一个重量的人的个数，第二次循环把每个人的重量在map里查找，相加等于limit的，到相加小于limit的，如果找到最优配对的，把两人从map里减一
// 最后把map里没有配对的人都单独装一艘船里
func numRescueBoats1(people []int, limit int) int {
	// 定义船的数量和map
	var boatNum int = 0
	var hsMap map[int]int = make(map[int]int)

	// 循环people，把每个人的重量放到map里
	for _, v := range people {
		if v < limit {
			hsMap[v] += 1
		}
	}

	// 循环map，找到最优配对的，把两人从map里减一
	for _, v := range people {
		// 如果有人的重量等于limit，直接加一艘船
		if v == limit {
			boatNum += 1
		} else {
			// 如果有人的重量小于limit，找到最优配对的，把两人从map里减一
			for i := limit - v; i > 0; i-- {
				// 如果两个人重量相等，而这个有这个重量的大于等于两人，直接加一艘船，把两人从map里减去
				if i == v {
					if hsMap[v] >= 2 {
						hsMap[v] -= 2
						boatNum += 1
						break
					}
				} else {
					// 如果两个人重量不等，找到最优配对的，把两人从map里减一，最优配对是两个人的重量相加等于limit，如果没有找到，就找到相加小于limit的
					if hsMap[i] > 0 && hsMap[v] > 0 {
						hsMap[i] -= 1
						hsMap[v] -= 1
						boatNum += 1
						break
					}
				}
			}
		}
	}

	// 循环map，把map里没有配对的人数都单独装一艘船里
	for _, v := range hsMap {
		boatNum += v
	}

	return boatNum
}

// =======================================================================================================================
// 2.双指针，排序 O(nlogn)-O(1)
// 先对people排序，然后用两个指针head和tail分别指向people的头和尾，
// 如果people[head]+people[tail] > limit，tail往前移动，如果people[head]+people[tail] <= limit，head往前移动，tail往前移动
func numRescueBoats(people []int, limit int) int {
	// 首先对people从小到大排序
	sort.Ints(people)
	// 定义船的数量和头尾指针
	var boatNum int = 0
	var head, tail int = 0, len(people) - 1

	// 循环people
	for head < tail {
		// 如果people[head]+people[tail] > limit，装重的那个人，tail往前移动，船的数量加一
		if people[head]+people[tail] > limit {
			boatNum += 1
			tail -= 1
			// 如果people[head]+people[tail] <= limit，轻重两人一起装，head往前移动，tail往前移动，船的数量加一
		} else if people[head]+people[tail] <= limit {
			boatNum += 1
			head += 1
			tail -= 1
		}
	}

	return boatNum
}

// =======================================================================================================================
// 3.双指针，优化版，不排序 O(n)-O(n)
// 用一个数组frequency记录每一个重量的人的个数，如果重量为i的人有n个，frequency[i] = n
func numRescueBoats2(people []int, limit int) int {

	// 定义一个数组frequency记录每一个重量的人的个数
	var frequency = make([]int, limit+1)
	for _, v := range people {
		frequency[v] += 1
	}

	// 定义船的数量和头尾指针
	var boatNum int = 0
	var head, tail int = 1, limit

	// 循环frequency
	for head <= tail {
		// 如果当前重量为head的人数为0，head+1，直到找到重量为head的人
		for head <= tail && frequency[head] == 0 {
			head += 1
		}
		// 如果当前重量为tail的人数为0，tail-1，直到找到重量为tail的人
		for head <= tail && frequency[tail] == 0 {
			tail -= 1
		}

		// 如果head > tail，说明所有人都装完了，跳出循环
		if head > tail {
			break
		}

		// 因为每个人的重量都小于等于limit，所以重的人一定上船，船的数量加一，数组frequency里的重的人数减一
		// 此时不知道轻的人是否上船，所以不移动head
		frequency[tail] -= 1
		boatNum += 1

		// 如果轻的人加上已经上船的重的人的重量小于等于limit，轻的人也可以上船，数组frequency里的轻的人数减一
		if head+tail <= limit && frequency[head] > 0 {
			frequency[head] -= 1
		}
	}

	// 返回船的数量
	return boatNum
}
```