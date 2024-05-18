package main

import "fmt"

func main() {
	str := "RDD"
	res := predictPartyVictory1(str)
	fmt.Println(res)
}

// =====================================================================================================================
// 1. 循环消除
func predictPartyVictory(senate string) string {
	arr := []byte(senate) // 把字符串转换为字节数组
	Damount := 0          // D的总数
	Ramount := 0          // R的总数
	Rvote := 0            // R的投票
	Dvote := 0            // D的投票

	// 统计D和R的总数
	for _, v := range arr {
		if v == 'D' {
			Damount += 1
		} else {
			Ramount += 1
		}
	}

	// 如果D或R的总数为0，直接宣布另外一方胜利
	if Damount == 0 {
		return "Radiant"
	} else if Ramount == 0 {
		return "Dire"
	}

	// 在数组中不断循环，直到有一方的总数为0，宣布另外一方胜利
	i := 0
	for {
		if arr[i] == 'R' { // 如果是R，当D的投票为0时，R的投票加1，否则D的投票减1，这个值也被设置为‘.’，R的总数减1
			if Dvote == 0 {
				Rvote += 1
			} else {
				Dvote -= 1
				arr[i] = '.'
				Ramount -= 1
				if Ramount == 0 { // 如果R的总数为0，宣布D胜利
					return "Dire"
				}
			}
		} else if arr[i] == 'D' { // 如果是D，当R的投票为0时，D的投票加1，否则R的投票减1，这个值也被设置为‘.’，D的总数减1
			if Rvote == 0 {
				Dvote += 1
			} else {
				Rvote -= 1
				arr[i] = '.'
				Damount -= 1
				if Damount == 0 { // 如果D的总数为0，宣布R胜利
					return "Radiant"
				}
			}
		}
		i += 1
		if i == len(arr) { // 如果i等于数组的长度，i重新赋值为0
			i = 0
		}
	}
}

// =====================================================================================================================
// 2. 双队列，行使ban权的参议院加入队尾。
func predictPartyVictory1(senate string) string {
	arr := []byte(senate)
	var listR, listD []int

	for i, v := range arr {
		if v == 'D' {
			listD = append(listD, i)
		} else {
			listR = append(listR, i)
		}
	}

	for len(listD) > 0 && len(listR) > 0 {

		// 这一步是关键，如果一个参议员行使权利，ban了另外一党的，它本身的顺位就加一轮，继续加入到队尾
		if listD[0] < listR[0] {
			listD = append(listD, listD[0]+len(senate))
		} else {
			listR = append(listR, listR[0]+len(senate))
		}

		listD = listD[1:]
		listR = listR[1:]
	}

	if len(listD) == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

// =====================================================================================================================
// 3. solution from leetcode
// 一个队列模拟,另外一个用队列或者栈储存，都可以
func predictPartyVictory2(senate string) string {
	l := len(senate)
	temp := make([]int, 0, l)
	queue := make([]int, 0, l)

	// 把字符串以01的方式存进队列
	for _, x := range senate {
		var y int
		if x == 'R' {
			y = 0
		} else {
			y = 1
		}
		queue = append(queue, y)
	}

	// 当队列不为零时，循环
	for len(queue) > 0 {
		// 取出队列第一个
		current := queue[0]
		queue = queue[1:]

		// 如果另一个（栈/队列）空或者栈最后一个与当前取出元素相同，加入栈
		if len(temp) == 0 || temp[len(temp)-1] == current {
			temp = append(temp, current)
		} else {
			// 否则得到另一个（栈/队列）第一个或最后一个，加入队列尾，丢弃当前取出元素
			popped := temp[0]
			temp = temp[1:]
			queue = append(queue, popped)
		}
	}

	// 最后当队列空时，另一个（栈/队列）里是什么元素，什么元素就赢
	if temp[len(temp)-1] == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

// =====================================================================================================================
