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
