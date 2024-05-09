package main

func main() {

}

// 1.设X，Y变量代表XY坐标轴，如果走完之后，X，Y都等于0，就是回到了原地
func judgeCircle(moves string) bool {
	var X, Y int = 0, 0
	for _, v := range moves {
		switch v {
		case 'U':
			Y += 1
		case 'D':
			Y -= 1
		case 'L':
			X -= 1
		case 'R':
			X += 1
		}
	}

	if X == 0 && Y == 0 {
		return true
	} else {
		return false
	}
}
