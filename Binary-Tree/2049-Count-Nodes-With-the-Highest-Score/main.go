package main

import "fmt"

func main() {
	parents := []int{-1, 2, 0}
	res := countHighestScoreNodes(parents)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo map[int]int //备忘录

// countHighestScoreNodes 计算具有最高分数的节点数量
func countHighestScoreNodes(parents []int) int {
	// 构建树的邻接表表示
	tree := build(parents)
	hsMap := make(map[int]int) // 存储每个分数出现的次数
	memo = make(map[int]int)   // 记录每个子树的节点数量
	highestScore := 0          // 当前最高分数

	for i := 0; i < len(parents); i++ {
		var score int = 1
		// 获取左右子树的节点数量，然后根据左右子树的节点数量计算当前节点的分数。
		leftNums, rightNums := getSubNums(tree, i)
		if leftNums != 0 {
			score *= leftNums
		}
		if rightNums != 0 {
			score *= rightNums
		}
		// 如果全部节点数减去左右子树的节点数和自身之后不为 0，则将其乘到分数中
		if len(parents)-leftNums-rightNums-1 != 0 {
			score *= len(parents) - leftNums - rightNums - 1
		}
		if score > highestScore {
			highestScore = score
		}
		hsMap[score] += 1
	}

	return hsMap[highestScore]
}

// build 构建树的邻接表
func build(parents []int) [][]int {
	n := len(parents)
	tree := make([][]int, n)
	for i := range tree {
		tree[i] = []int{-1, -1}
	}

	// 获取每个节点的子节点
	for i := 1; i < n; i++ {
		if tree[parents[i]][0] == -1 {
			tree[parents[i]][0] = i
		} else {
			tree[parents[i]][1] = i
		}
	}

	return tree
}

// getSubNums 获取节点 i 的左右子树的节点数量
func getSubNums(tree [][]int, i int) (int, int) {
	if i >= len(tree) {
		return 0, 0
	}

	// 计算左右子树的节点数量
	var sub1Num, sub2Num int
	// 如果子树的节点数量已经计算过，则直接返回
	if _, exist := memo[tree[i][0]]; exist {
		sub1Num = memo[tree[i][0]]
	} else {
		// 如果没有计算过，则递归计算，并将结果存入备忘录
		sub1Num = dfs(tree, tree[i][0])
		memo[tree[i][0]] = sub1Num
	}

	if _, exist := memo[tree[i][1]]; exist {
		sub2Num = memo[tree[i][1]]
	} else {
		sub2Num = dfs(tree, tree[i][1])
		memo[tree[i][1]] = sub2Num
	}

	return sub1Num, sub2Num
}

// dfs 计算节点 i 的子树的节点数量
func dfs(tree [][]int, i int) int {
	if i == -1 {
		return 0
	}

	// 如果已经计算过，则直接返回
	if _, exist := memo[i]; exist {
		return memo[i]
	}

	return dfs(tree, tree[i][0]) + dfs(tree, tree[i][1]) + 1
}
