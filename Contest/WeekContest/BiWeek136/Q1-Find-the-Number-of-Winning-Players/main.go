package main

func main() {

}

func winningPlayerCount(n int, pick [][]int) int {
	player := make([][]int, n)
	for i := range player {
		player[i] = make([]int, 11)
	}
	winner := make([]int, n)
	for _, v := range pick {
		player[v[0]][v[1]] += 1
		if player[v[0]][v[1]] > v[0] {
			winner[v[0]] = 1
		}
	}

	res := 0
	for _, v := range winner {
		if v == 1 {
			res += 1
		}
	}

	return res
}
