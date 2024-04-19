# 1672. Richest Customer Wealth

Easy

You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

 

Example 1:

> Input: accounts = \[[1,2,3],[3,2,1]]
Output: 6
Explanation:
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.

Example 2:

>Input: accounts = \[[1,5],[7,3],[3,5]]
Output: 10
Explanation: 
1st customer has wealth = 6
2nd customer has wealth = 10 
3rd customer has wealth = 8
The 2nd customer is the richest with a wealth of 10.

Example 3:

>Input: accounts = \[[2,8,7],[7,1,3],[1,9,5]]
Output: 17
 

Constraints:

m == accounts.length
n == accounts[i].length
1 <= m, n <= 50
1 <= accounts[i][j] <= 100

## My First Try (√)

```go
func maximumWealth(accounts [][]int) int {
	// 假设最富有的投资者资金为0
	var oneManInv int = 0

	// 循环每个人，得到单个人在每个银行的投资切片[]int
	for _, val := range accounts {
		// 此人在每个银行总共投资总数为sum
		var sum int
		// 把每个银行的钱加起来
		for _, val := range val {
			sum += val
		}

		//如果总数大于假设值，则这个人成为新的最富有者，金额为oneManInv
		if oneManInv < sum {
			oneManInv = sum
		}
	}

	//返回最富投资者资金总数
	return oneManInv
}
```

## My Second Try (√) -- Concurrency

```go
func maximumWealth(accounts [][]int) int {
	// 假设最富有的投资者资金为0
	sumMax := 0
	// 定义一个任务数量为len(accounts)，也就是统计的人数的通道
	var sumChan chan int = make(chan int, len(accounts))
	// 并发计算每个人的资金总和
	for _, arr := range accounts {
		go sum(arr, sumChan)
	}

	// 如果发现有高于现有sumMax的资金，将其替换为sumMax
	for i := 0; i < len(accounts); i++ {
		newSum := <-sumChan // 从通道读取一次，保存到变量中
		if sumMax < newSum {
			sumMax = newSum
		}
	}

	// 返回最大资金
	return sumMax
}

// 定义一个函数，计算一个整数切片的所有数之和
// 接受一个整数型切片和一个通道，结果返回给通道
func sum(arr []int, result chan<- int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	result <- sum
}
```

使用 goroutines 和 channels 来并发计算每个客户的财富看似提升了效率，但对于这种计算密集型但计算规模较小的任务，引入并发可能反而增加了复杂度和执行时间（因为goroutine的调度和通道的同步操作本身有开销）。对于小规模数据，直接串行处理通常更快，也更简单。