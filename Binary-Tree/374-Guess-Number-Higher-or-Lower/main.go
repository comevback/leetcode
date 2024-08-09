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
