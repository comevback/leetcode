package main

import "fmt"

func main() {
	heights := []int{10, 6, 8, 5, 11, 9}
	res := canSeePersonsCount(heights)
	fmt.Println(res)
}

// canSeePersonsCount 接受一个整数切片 heights，每个整数代表一个人的高度。
// 函数返回一个整数切片，每个元素表示每个人可以看到的后面人数。
func canSeePersonsCount(heights []int) []int {
	res := make([]int, len(heights)) // 初始化结果切片，长度与输入的高度相同
	temp := []int{}                  // 用作栈，存储从右向左处理过程中的人的高度

	// 从右向左遍历高度数组
	for i := len(heights) - 1; i >= 0; i-- {
		nums := 0 // 初始化计数器，记录当前人可以看到的人数
		// 循环直到栈为空或者栈顶的高度大于等于当前人的高度
		for len(temp) > 0 && heights[i] > temp[len(temp)-1] {
			nums += 1                 // 如果栈顶的人更矮，当前人可以看到这个人
			temp = temp[:len(temp)-1] // 移除栈顶元素，因为这个矮人已经被当前的高人看到了
		}
		// 如果栈不为空，说明栈顶的人比当前人更高
		if len(temp) != 0 {
			nums += 1 // 当前人还可以看到栈顶的这个更高的人
		}
		res[i] = nums                   // 将当前人可以看到的人数存入结果数组
		temp = append(temp, heights[i]) // 将当前人的高度推入栈中，为之后的人提供比较基准
	}

	return res // 返回结果数组
}
