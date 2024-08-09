# 374. Guess Number Higher or Lower

Solved
Easy
Topics
Companies
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Example 1:

Input: n = 10, pick = 6
Output: 6
Example 2:

Input: n = 1, pick = 1
Output: 1
Example 3:

Input: n = 2, pick = 1
Output: 1

Constraints:

1 <= n <= 231 - 1
1 <= pick <= n

---

# Code

```go
package main

func main() {

}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    // 二分查找
	left, right := 1, n
	mid := left + (right-left)/2

    // 当mid在1到n之间时，进行二分查找
	for mid >= 1 && mid <= n {
		mid = left + (right-left)/2
		temp := guess(mid)
		if temp == 0 { // 如果猜中了，返回mid，否则根据temp的值调整left和right
			return mid
		} else if temp == 1 {
			left = mid + 1
		} else if temp == -1 {
			right = mid - 1
		}
	}

    // 最后返回left
	return left
}
```
