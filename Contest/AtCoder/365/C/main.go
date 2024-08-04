package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// 读取输入的花费数量和最大允许花费
	var nums, maxMoney int
	fmt.Scanf("%d %d", &nums, &maxMoney)

	// 设置读取输入的扫描器
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	expens := []int{}
	sum := 0
	max := math.MinInt

	// 读取所有花费
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		if num > max {
			max = num
		}
		sum += num
		expens = append(expens, num)

		nums -= 1
		if nums < 1 {
			break
		}
	}

	// 如果总花费小于或等于最大允许花费，输出 "infinite"
	if sum <= maxMoney {
		fmt.Println("infinite")
		return
	}

	// 需要减少的总花费
	diff := sum - maxMoney

	// 二分查找确定最大临界值
	left, right := 0, max
	res := -1

	for left <= right {
		temp := 0
		mid := left + (right-left)/2
		// 计算调整后的总花费
		for _, v := range expens {
			if v > mid {
				temp += v - mid
			}
		}
		// 如果找到恰好满足条件的值，直接返回
		if diff == temp {
			fmt.Println(mid)
			return
		} else if diff < temp {
			// 如果调整后的总花费大于需要减少的花费，尝试更大的临界值
			res = mid
			left = mid + 1
		} else {
			// 如果调整后的总花费小于需要减少的花费，尝试更小的临界值
			right = mid - 1
		}
	}

	// 输出结果
	fmt.Println(res)
}
