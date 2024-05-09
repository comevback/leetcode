package main

import "fmt"

func main() {
	var moves [][]int = [][]int{{2, 2}, {0, 1}, {2, 1}, {0, 0}, {0, 2}}
	fmt.Println(tictactoe(moves))
}

// 1. 二维数组模拟棋盘，找一条线上元素是否相同且不为空
func tictactoe(moves [][]int) string {
	// 如果小于5步，返回Pending
	if len(moves) < 5 {
		return "Pending"
	}
	// 棋盘
	var chess [3][3]string
	// i, j分别代表A和B的步数，i从0开始，j从1开始
	i, j := 0, 1
	// 遍历moves
	for i < len(moves) || j < len(moves) {

		// 如果i小于len(moves)，则在chess[moves[i][0]][moves[i][1]]处放入“A”
		if i < len(moves) {
			chess[moves[i][0]][moves[i][1]] = "A"
			i += 2
		}

		// 如果j小于len(moves)，则在chess[moves[j][0]][moves[j][1]]处放入“B”
		if j < len(moves) {
			chess[moves[j][0]][moves[j][1]] = "B"
			j += 2
		}
	}

	// 遍历棋盘，横竖方向是否有一条线上元素相同且不为空，如果有，返回该元素
	for i := 0; i < 3; i++ {
		if chess[i][0] == chess[i][1] && chess[i][1] == chess[i][2] && chess[i][0] != "" {
			return chess[i][0]
		} else if chess[0][i] == chess[1][i] && chess[1][i] == chess[2][i] && chess[0][i] != "" {
			return chess[0][i]
		}
	}

	// 遍历棋盘，对角线方向是否有一条线上元素相同且不为空，如果有，返回该元素
	if chess[0][0] == chess[1][1] && chess[1][1] == chess[2][2] && chess[1][1] != "" {
		return chess[0][0]
	} else if chess[2][0] == chess[1][1] && chess[1][1] == chess[0][2] && chess[1][1] != "" {
		return chess[2][0]
	}

	// 如果以上条件都不满足，如果moves长度小于9，返回Pending，否则返回Draw
	if len(moves) < 9 {
		return "Pending"
	} else {
		return "Draw"
	}
}
