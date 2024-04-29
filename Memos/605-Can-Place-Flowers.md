# 605. Can Place Flowers

Easy

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

 

Example 1:
> Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
> Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
 

Constraints:
> 1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length

---

```go
package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 1, 0, 0, 1, 0}
	n := 2
	res := canPlaceFlowers1(flowerbed, n)
	fmt.Println(res)
	flowerbed = []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0}
	n = 17
	res = canPlaceFlowers1(flowerbed, n)
	fmt.Println(res)
}

// 1.从前往后挨个检查，是否满足条件
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	// 如果只有一个元素，且为0，或者为1且n为0，都可以种植，返回true，否则返回false
	if length == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		if flowerbed[0] == 1 && n == 0 {
			return true
		} else {
			return false
		}
	}

	// 如果不止一个元素，设可以种植的植物数为possibleNum
	var possibleNum int = 0

	// 从前往后挨个检查，是否满足条件
	for i := 0; i < length; i++ {
		// 必须元素当前值为0，且前后位置都为0，才可以种植
		if flowerbed[i] == 0 {
			// 如果是第一个元素，且后一个元素为0，可以种植
			if i == 0 {
				if flowerbed[1] == 0 {
					flowerbed[0] = 1
					possibleNum += 1
				}
				continue
			}
			// 如果是最后一个元素，且前一个元素为0，可以种植
			if i == length-1 {
				if flowerbed[length-2] == 0 {
					flowerbed[length-1] = 1
					possibleNum += 1
				}
				continue
			}
			// 如果是中间元素，而且前后元素都为0，可以种植
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				possibleNum += 1
			}
		}
	}

	// 判断是否可以种植的植物数是否大于等于n，是则返回true，否则返回false
	if possibleNum >= n {
		return true
	} else {
		return false
	}
}

// ================================================================================================================================
// 2.与上解法一样
func canPlaceFlowers1(flowerbed []int, n int) bool {
	length := len(flowerbed)
	// 如果只有一个元素，且为0，或者为1且n为0，都可以种植，返回true，否则返回false
	if length == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		if flowerbed[0] == 1 {
			return false
		}
	}

	if n == 0 {
		return true
	}

	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			// 如果是第一个元素，且后一个元素为0，可以种植
			if i == 0 {
				if flowerbed[1] == 0 {
					flowerbed[0] = 1
					n -= 1
					if n == 0 {
						return true
					}
				}
				continue
			}
			// 如果是最后一个元素，且前一个元素为0，可以种植
			if i == length-1 {
				if flowerbed[length-2] == 0 {
					flowerbed[length-1] = 1
					n -= 1
					if n == 0 {
						return true
					}
				}
				continue
			}
			if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n -= 1
				if n == 0 {
					return true
				}
			}
		}
	}

	return false
}

// 与2一样，每次遇到能种的位置就把n减一，但在循环中如果种了花，则跳过下一个位置
func canPlaceFlowers3(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		// 首先这个位置必须为空
		if flowerbed[i] == 0 &&
			// 然后要么是最后一个，要么后一位为空
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) &&
			// 要么是第一个，要么前一位为空
			(i == 0 || flowerbed[i-1] == 0) {
			// 满足三个条件的，就是可种植的
			n -= 1
			// 跳过下一个位置，就不用考虑之前一个位置是否种花了
			i += 1
		}
	}

	return n <= 0
}

// ================================================================================================================================
// 3.我希望算出整个花坛最多能装多少花，然后遍历累加已有花数量，加上n，与最大能装花数量对比（X）
func canPlaceFlowers2(flowerbed []int, n int) bool {
	length := len(flowerbed)
	// 如果只有一个元素，且为0，或者为1且n为0，都可以种植，返回true，否则返回false
	if length == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		if flowerbed[0] == 1 && n == 0 {
			return true
		} else {
			return false
		}
	}

	var mostFlowerNum int = length / 2
	if length%2 == 0 && flowerbed[0] == 1 && flowerbed[length-1] == 1 {
		mostFlowerNum = length/2 - 1
	}
	if length%2 == 1 {
		mostFlowerNum = length/2 + 1
	}

	flowerNum := n
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			flowerNum += 1
		}
		if length%2 == 1 {
			if (i+1)%2 == 0 && flowerbed[i] == 1 {
				mostFlowerNum = length / 2
			}
		}
	}

	if length%2 == 1 && flowerNum == n {
		mostFlowerNum = length/2 + 1
	}

	if flowerNum > mostFlowerNum {
		return false
	} else {
		return true
	}
}

```