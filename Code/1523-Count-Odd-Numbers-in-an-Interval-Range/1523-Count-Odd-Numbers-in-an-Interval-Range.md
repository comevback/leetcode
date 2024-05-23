# Count Odd Numbers in an Interval Range

Category	Difficulty	Likes	Dislikes
algorithms	Easy (49.88%)	2737	157

Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

 

Example 1:
> Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].

Example 2:
> Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].
 

Constraints:
> 0 <= low <= high <= 10^9
Submissions | Solution

---

# Code
```go
func countOdds(low int, high int) int {
    var count int
	if low%2 == 1 || high%2 == 1 {
		count = (high-low)/2 + 1
	} else {
		count = (high - low) / 2
	}
	return count
}
```

# Code
```go
func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}
```

计算公式：
奇数数量 = (high + 1) / 2 - low / 2
1. 为什么要用 (high + 1) / 2：
这个表达式用来找出从 1 到 high 之间包含多少个奇数。举个例子：
如果 high 是偶数，如 10，那么奇数是 1, 3, 5, 7, 9，一共有 5 个。(10 + 1) / 2 = 5.5，取整为 5。
如果 high 是奇数，如 11，那么奇数是 1, 3, 5, 7, 9, 11，一共有 6 个。(11 + 1) / 2 = 6。
所以这个表达式基本上计算的是从 1 到 high 的序列中奇数的总数。

2. 为什么要用 low / 2：
这个表达式用来找出从 1 到 low-1 之间包含多少个奇数。比如：
如果 low 是 5，那么从 1 到 4 的奇数有 1, 3，一共 2 个奇数。5 / 2 = 2.5，取整为 2。
这样，通过这个表达式，我们可以计算出在 low 之前（不包括 low）有多少个奇数。

3. 计算区间内奇数的数量：
通过以上两步，我们可以计算出 [1, high] 区间内的奇数数量减去 [1, low-1] 区间内的奇数数量，就得到了 [low, high] 区间内的奇数数量。

