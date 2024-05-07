# 1732. Find the Highest Altitude

Easy

Hint
There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.


Example 1:
> Input: gain = [-5,1,5,0,-7]
Output: 1
Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.

Example 2:
> Input: gain = [-4,-3,-2,-1,4,3,2]
Output: 0
Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
 

Constraints:
> n == gain.length
1 <= n <= 100
-100 <= gain[i] <= 100

---

# Code
```go
package main

func main() {

}

// 1.一次遍历
func largestAltitude(gain []int) int {
	// 设海拔为alti
	alti := 0
	// 达到的最高海拔为maxAlti
	maxAlti := 0

	// 遍历gain，累加海拔，如果当前海拔大于最高海拔，替换
	for _, v := range gain {
		alti += v
		if alti > maxAlti {
			maxAlti = alti
		}
	}

	// 返回最高海拔
	return maxAlti
}
```