package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var T, B, U, D, L, R int
	var step int
	fmt.Scanf("%d %d %d %d %d %d", &T, &B, &U, &D, &L, &R)
	fmt.Scan(&step)
	var current int
	current = T
	allRotateTimes := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		if (current == T && target == B) || (current == B && target == T) || (current == U && target == D) || (current == D && target == U) || (current == L && target == R) || (current == R && target == L) {
			allRotateTimes += 2
		} else if current == target {
			allRotateTimes += 0
		} else {
			allRotateTimes += 1
		}

		current = target

		step -= 1
		if step < 1 {
			break
		}
	}

	fmt.Println(allRotateTimes)
}
