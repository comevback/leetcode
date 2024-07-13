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
