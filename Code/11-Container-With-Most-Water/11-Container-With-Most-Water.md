# 11. Container With Most Water（双指针）

Medium

Hint
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

 

Example 1:
![11-Container-With-Most-Water/11-Container-With-Most-Water.png](11-Container-With-Most-Water.png)
> Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
> Input: height = [1,1]
Output: 1
 

Constraints:
> n == height.length
2 <= n <= 105
0 <= height[i] <= 104

---

## 思路

> 双指针 O(n)-O(1)
设一头一尾两个指针，当前面积取决于两个指针的距离和两个指针的最小值：amount := (tail - head) * min(height[head], height[tail])
每次移动较小的那个指针，直到两个指针相遇，过程中记录更新最大面积

## 代码

```go
package main

import "fmt"

func main() {
	var height []int = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

// 1.暴力循环 O(n^2)-O(1) 错误
// 把每个长度的都循环一遍，找出最大值
func maxArea1(height []int) int {
	// var head, tail int = 0, len(height) - 1
	var maxWater int = 0

	for i := 1; i < len(height); i++ {
		for j := 0; j < len(height)-i; j++ {
			if i*min(height[j], height[j+i]) > maxWater {
				maxWater = i * min(height[j], height[j+i])
			}
		}
	}

	return maxWater

}

// ========================================================================================================================
// 2.双指针 O(n)-O(1)
// 设一头一尾两个指针，当前面积取决于两个指针的距离和两个指针的最小值：amount := (tail - head) * min(height[head], height[tail])
// 每次移动较小的那个指针，直到两个指针相遇，过程中记录更新最大面积
func maxArea(height []int) int {
	// 设一头一尾两个指针
	var head, tail int = 0, len(height) - 1
	// 初始化面积为：头尾两个指针的距离 * 两个指针的最小值
	var maxWater int = (tail - head) * min(height[head], height[tail])

	// 循环移动两个指针
	for head < tail {
		// 移动较小的那个指针
		if height[head] < height[tail] {
			head += 1
		} else {
			tail -= 1
		}
		// 计算当前面积
		amount := (tail - head) * min(height[head], height[tail])
		// 如果当前面积比最大面积更大，更新最大面积
		if amount > maxWater {
			maxWater = amount
		}
	}

	// 返回最大面积
	return maxWater
}
```