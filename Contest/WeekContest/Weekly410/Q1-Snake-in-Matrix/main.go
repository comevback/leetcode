package main

func main() {

}

func finalPositionOfSnake(n int, commands []string) int {
	startPos := []int{0, 0}

	for _, v := range commands {
		switch v {
		case "UP":
			startPos[0] -= 1
		case "DOWN":
			startPos[0] += 1
		case "LEFT":
			startPos[1] -= 1
		case "RIGHT":
			startPos[1] += 1
		}
	}

	res := startPos[0]*n + startPos[1]
	return res
}
