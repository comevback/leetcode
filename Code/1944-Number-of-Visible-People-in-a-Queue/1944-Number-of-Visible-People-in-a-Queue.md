# 1944. Number of Visible People in a Queue

Solved
Hard
Topics
Companies
Hint
There are n people standing in a queue, and they numbered from 0 to n - 1 in left to right order. You are given an array heights of distinct integers where heights[i] represents the height of the ith person.

A person can see another person to their right in the queue if everybody in between is shorter than both of them. More formally, the ith person can see the jth person if i < j and min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]).

Return an array answer of length n where answer[i] is the number of people the ith person can see to their right in the queue.

Example 1:

Input: heights = [10,6,8,5,11,9]
Output: [3,1,2,1,1,0]
Explanation:
Person 0 can see person 1, 2, and 4.
Person 1 can see person 2.
Person 2 can see person 3 and 4.
Person 3 can see person 4.
Person 4 can see person 5.
Person 5 can see no one since nobody is to the right of them.
Example 2:

Input: heights = [5,1,2,3,10]
Output: [4,1,1,1,0]

Constraints:

n == heights.length
1 <= n <= 105
1 <= heights[i] <= 105
All the values of heights are unique.

---

# Code

```go
package main

import "fmt"

func main() {
	heights := []int{10, 6, 8, 5, 11, 9}
	res := canSeePersonsCount(heights)
	fmt.Println(res)
}

// canSeePersonsCount 接受一个整数切片 heights，每个整数代表一个人的高度。
// 函数返回一个整数切片，每个元素表示每个人可以看到的后面人数。
func canSeePersonsCount(heights []int) []int {
	res := make([]int, len(heights)) // 初始化结果切片，长度与输入的高度相同
	temp := []int{}                  // 用作栈，存储从右向左处理过程中的人的高度

	// 从右向左遍历高度数组
	for i := len(heights) - 1; i >= 0; i-- {
		nums := 0 // 初始化计数器，记录当前人可以看到的人数
		// 循环直到栈为空或者栈顶的高度大于等于当前人的高度
		for len(temp) > 0 && heights[i] > temp[len(temp)-1] {
			nums += 1                 // 如果栈顶的人更矮，当前人可以看到这个人
			temp = temp[:len(temp)-1] // 移除栈顶元素，因为这个矮人已经被当前的高人看到了
		}
		// 如果栈不为空，说明栈顶的人比当前人更高
		if len(temp) != 0 {
			nums += 1 // 当前人还可以看到栈顶的这个更高的人
		}
		res[i] = nums                   // 将当前人可以看到的人数存入结果数组
		temp = append(temp, heights[i]) // 将当前人的高度推入栈中，为之后的人提供比较基准
	}

	return res // 返回结果数组
}
```
