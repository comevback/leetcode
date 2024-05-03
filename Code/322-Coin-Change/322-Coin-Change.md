# 322. Coin Change（动态规划）（背包，零钱问题）

Medium

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
> Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
> Input: coins = [2], amount = 3
Output: -1

Example 3:
> Input: coins = [1], amount = 0
Output: 0
 

Constraints:
> 1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

---

## 1. Dynamic Programming
My code:
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

// Constraints:

// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2^31 - 1
// 0 <= amount <= 10^4

// ==========================================================================================================================
// 动态规划方法。具体思路如下：

// 定义状态： 创建一个数组 dp，其中 dp[i] 表示凑成总金额 i 所需的最少硬币数。
// 初始化： 初始化 dp[0] 为 0，因为总金额为 0 时不需要任何硬币。对于其他所有金额，可以先初始化为一个较大值（例如 amount + 1），表示无法凑成该金额。
// 状态转移方程： 对于每一个金额 i 从 1 到 amount，遍历每一个硬币面额 coin。如果 coin 小于或等于 i，因为只有可能用比amount小的数额钞票，则有：

//                    dp[i]=min(dp[i],dp[i−coin]+1)

// 这里的意思是，如果你考虑使用面额为 coin 的硬币，那么需要的硬币数量将是凑成金额 i - coin 所需的硬币数加上这一个硬币。
// 结果： 计算完所有 dp 值后，检查 dp[amount] 的值。如果它仍然是初始的较大值，表示无法凑出总金额，返回 -1。否则，返回 dp[amount]，即为所需的最少硬币数。
// 这种方法的时间复杂度是 O(n \times amount)，其中 n 是硬币种类的数量，空间复杂度是 O(amount)，因为需要一个大小为 amount + 1 的数组来存储中间结果。
// ==========================================================================================================================

// 1.动态规划
// 状态定义：dp[i]表示凑成总金额i所需的最少硬币数
// 要求某一个金额i的最少硬币数，可以考虑遍历每一个硬币面额coin，如果coin小于等于i，则可能使用这个硬币，再加上dp[i-coin]即可得到凑成金额i所需的硬币数
// 举例来说，假设coins=[1,2,5]，amount=11，那么dp[11]可以由dp[10]、dp[9]、dp[6]推导而来，因为11可以由10+1、9+2、6+5得到。
func coinChange(coins []int, amount int) int {
	// 初始化dp数组，dp[i]表示凑成总金额i所需的最少硬币数
	var dp []int = make([]int, amount+1)
	// 把数组每一项初始化为-1，如果无法凑成总金额i，则dp[i]为-1
	for i := range dp {
		dp[i] = -1
	}
	// 初始化dp[0]为0，因为凑成总金额0不需要任何硬币
	dp[0] = 0

	// 遍历每一个金额i从1到amount
	for i := 1; i < amount+1; i++ {
		// 初始化一个变量fewest，用来记录凑成金额i所需的最少硬币数，初始化为int的最大值
		var fewest int = math.MaxInt32
		// 遍历每一个硬币面额coin，如果coin小于等于i，则可能使用这个硬币
		for _, coin := range coins {
			// 如果coin大于i，则无法使用这个硬币
			if coin > i {
				continue
			}
			// 如果coin等于i，则凑成金额i所需的最少硬币数为1
			if i == coin {
				dp[i] = 1
				break
			}
			// 如果coin小于i，则可能使用这个硬币，再加上dp[i-coin]即可得到凑成金额i所需的硬币数，所以凑成金额i所需的最少硬币数为dp[i-coin]+1
			// 如果dp[i-coin]为-1，则表示无法凑成金额i-coin，所以无法凑成金额i
			if i > coin && dp[i-coin] < fewest && dp[i-coin] != -1 {
				fewest = dp[i-coin]
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	// 上面的循环可以简化为如下代码
	// for i := 1; i <= amount; i++ {
	// 	for _, coin := range coins {
	// 		if i >= coin {
	// 			dp[i] = min(dp[i], dp[i-coin]+1)
	// 		}
	// 	}
	// }

	// 返回dp[amount]，即为所需的最少硬币数
	return dp[amount]
}
```

## 2. Dynamic Programming (Optimized) by ChatGPT
```go
您的动态规划方法基本正确，但存在一些可以优化的地方以提升代码的效率和清晰度。以下是针对您的代码的一些建议和改进：

### 1. 初始值设定问题
您初始化 `dp` 数组时使用了 `-1` 来表示无法凑成该金额，这是一个合理的选择。然而，在初始化时，您可以只将 `dp[0]` 设置为 `0`，其他值可以初始化为一个很大的数，比如 `amount + 1`，这样可以避免在后续操作中频繁检查 `-1` 的条件。

### 2. 不必要的条件检查
在遍历硬币面额时，您检查了 `coin == i` 的情况并单独处理。其实这种情况可以不用单独处理，因为当 `coin == i` 时，`dp[i - coin]` 就是 `dp[0]`，也就是 `0`，因此 `dp[i]` 可以直接设置为 `1`，这在后续的计算中自然会得到处理。

### 3. 优化遍历硬币的逻辑
在您的代码中，您使用了一个 `fewest` 变量来保存最小硬币数，实际上可以直接在每次内部循环中更新 `dp[i]`，减少一些中间变量的使用。

### 4. 更简洁的动态规划逻辑
您可以直接在遍历硬币的过程中更新 `dp[i]`，不必使用额外的 `fewest` 变量。这样代码会更加简洁和直观。

下面是根据以上建议重写的代码：

```go
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1  // 使用一个大于amount的值初始化，便于后续比较
	}
	dp[0] = 0  // 凑成总金额0需要0个硬币

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount { // 如果dp[amount]仍然是初始化的值，说明无法凑出
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("Minimum coins required:", coinChange(coins, amount))
}
```

这个改进版的代码更加简洁，且通过直接在动态规划数组 `dp` 中更新最小值，减少了额外的比较和变量使用。此外，通过判断 `dp[amount]` 是否大于 `amount` 来确定是否有解，这是因为在没有合适硬币组合的情况下，`dp[amount]` 的值不会被更新成小于或等于 `amount` 的值。