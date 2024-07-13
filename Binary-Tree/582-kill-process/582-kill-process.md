# 582. 杀掉进程

系统中存在 n 个进程，形成一个有根树结构。给你两个整数数组 pid 和 ppid ，其中 pid[i] 是第 i 个进程的 ID ，ppid[i] 是第 i 个进程的父进程 ID 。

每一个进程只有 一个父进程 ，但是可能会有 一个或者多个子进程 。只有一个进程的 ppid[i] = 0 ，意味着这个进程 没有父进程 。

当一个进程 被杀掉 的时候，它所有的子进程和后代进程都要被杀掉。

给你一个整数 kill 表示要杀掉 ​​ 进程的 ID ，返回被杀掉的进程的 ID 列表。可以按 任意顺序 返回答案。

示例 1：

输入：pid = [1,3,10,5], ppid = [3,0,5,3], kill = 5
输出：[5,10]
解释：涂为红色的进程是应该被杀掉的进程。
示例 2：

输入：pid = [1], ppid = [0], kill = 1
输出：[1]
提示：

n == pid.length
n == ppid.length
1 <= n <= 5 _ 104
1 <= pid[i] <= 5 _ 104
0 <= ppid[i] <= 5 \* 104
仅有一个进程没有父进程
pid 中的所有值 互不相同
题目数据保证 kill 在 pid 中

---

# Code

```go
package main

import "fmt"

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	res := killProcess(pid, ppid, kill)
	fmt.Println(res)
}

func killProcess(pid []int, ppid []int, kill int) []int {
	// 建立一个哈希表，将每个父进程对应的子进程列表存储起来
	hsMap := make(map[int][]int)
	for i := 0; i < len(pid); i++ {
		hsMap[ppid[i]] = append(hsMap[ppid[i]], pid[i])
	}

	// 初始化结果数组
	res := []int{}
	// 初始化队列，将需要杀死的进程ID加入队列
	queue := []int{}
	queue = append(queue, kill)

	// 进行广度优先搜索（BFS）
	for len(queue) != 0 {
		// 当前队列的长度
		length := len(queue)
		for i := 0; i < length; i++ {
			// 将当前进程加入结果数组
			process := queue[i]
			res = append(res, process)
			// 将当前进程的所有子进程加入队列
			queue = append(queue, hsMap[process]...)
		}
		// 更新队列，去除已处理的进程
		queue = queue[length:]
	}
	return res
}
```
