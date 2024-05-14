# 735. Asteroid Collision

Medium

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Example 1:
> Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Example 2:
> Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.

Example 3:
> Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
 

Constraints:
> 2 <= asteroids.length <= 104
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0

---

# Code
```go
package main

func main() {

}

// 1. 用数组实现栈
func asteroidCollision(asteroids []int) []int {
	// 定义栈数组
	var stack []int
	// 定义asteroids里的指针
	current := 0

	// 当指针没遍历完asteroids时继续循环
	for current < len(asteroids) {
		// 如果栈空，或者新遇到的小行星方向相同，或者前左后右，就把新星加入栈
		// (stack[len(stack)-1] < 0 && asteroids[current] > 0) 是因为负数是向左的小行星，正数向右，如果现有向左的在栈中，遇到向右的，不会触发相撞
		if len(stack) == 0 || stack[len(stack)-1]*asteroids[current] > 0 || (stack[len(stack)-1] < 0 && asteroids[current] > 0) {
			stack = append(stack, asteroids[current])
			current += 1
		} else {
			// 如果是前右后左，就比较绝对值大小，如果前面大，直接指向下一个新星
			if abs(stack[len(stack)-1]) > abs(asteroids[current]) {
				current += 1
				// 如果后面大，把栈顶星去除
			} else if abs(stack[len(stack)-1]) < abs(asteroids[current]) {
				stack = stack[:len(stack)-1]
				// 如果一样大，两个相抵消
			} else {
				stack = stack[:len(stack)-1]
				current += 1
			}
		}
	}

	// 最后返回栈，就是结果
	return stack
}

// 求数的绝对值方法
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
```