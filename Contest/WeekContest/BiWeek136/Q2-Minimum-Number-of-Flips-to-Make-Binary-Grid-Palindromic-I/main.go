package main

func main() {

}

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowNeeds := 0
	colNeeds := 0
	for _, v := range grid {
		rowNeeds += makePali(v)
	}

	for i := 0; i < n; i++ {
		head, tail := 0, m-1
		res := 0
		for head < tail {
			if grid[head][i] != grid[tail][i] {
				res += 1
			}
			head += 1
			tail -= 1
		}
		colNeeds += res
	}

	return min(rowNeeds, colNeeds)
}

func makePali(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	res := 0

	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[head] != nums[tail] {
			res += 1
		}
		head += 1
		tail -= 1
	}

	return res
}

func min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	} else {
		return num2
	}
}
