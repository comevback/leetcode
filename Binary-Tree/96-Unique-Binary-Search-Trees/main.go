package main

func main() {
	n := 18
	res := numTrees(n)
	println(res)
}

// 用数组作为备忘录
// var memo [][]int = [][]int{}
// 用map作为备忘录
var memoMap map[[2]int]int

func numTrees(n int) int {
	// 初始化数组备忘录
	// memo = make([][]int, n+1)
	// for i := 1; i <= n; i++ {
	//     memo[i] = make([]int, n+1)
	// }

	// 初始化map备忘录
	memoMap = make(map[[2]int]int)

	res := build(1, n)
	return res
}

// 计算闭区间 [lo, hi] 的答案
func build(lo int, hi int) int {
	// 如果闭区间为空，返回 1
	if lo > hi {
		return 1
	}

	// 如果结果已经计算过，直接返回
	// if memo[lo][hi] != 0 {
	//     return memo[lo][hi]
	// }

	// 如果结果已经计算过，直接返回
	if memoMap[[2]int{lo, hi}] != 0 {
		return memoMap[[2]int{lo, hi}]
	}

	// 尝试每个数字作为根节点
	res := 0
	for i := lo; i <= hi; i++ {
		leftNums := build(lo, i-1)
		rightNums := build(i+1, hi)

		res += leftNums * rightNums
	}

	// 将结果存入备忘录
	// memo[lo][hi] = res
	memoMap[[2]int{lo, hi}] = res

	// 返回闭区间 [lo, hi] 的答案
	return res
}
