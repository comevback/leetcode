package main

func main() {

}

// 1. 从高位到地位，优先取反
// 遍历每个数组，如果第一个数是0，就把这个数组取反，因为第一位是最高位，取反后一定比不取反的大
// 然后遍历每一列，如果0的个数大于一半，就把这一列取反
func matrixScore(grid [][]int) int {
	// 设定一个变量sum，用来存储最后的结果
	sum := 0
	// 遍历每个数组，如果第一个数是0，就把这个数组取反
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			toggleRow(grid[i])
		}
	}

	// 遍历每一列，如果0的个数大于一半，就把这一列取反
	for j := 0; j < len(grid[0]); j++ {
		num := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 0 {
				num += 1
				if num > len(grid)/2 {
					toggleColumn(grid, j)
					break
				}
			}
		}
	}

	// 计算最后的结果
	for i := 0; i < len(grid); i++ {
		sum += arrTo1B(grid[i])
	}
	// 返回结果
	return sum
}

// 把数组转换为二进制数
func arrTo1B(nums []int) int {
	bNum := 0
	for _, v := range nums {
		bNum = bNum << 1 // 把二进制数左移一位
		bNum |= v        // 把当前数加到二进制数的最后一位
	}
	return bNum // 返回二进制数
}

// 把一个数组的所有数取反
func toggleRow(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}
}

// 把一个二维数组的某一列取反
func toggleColumn(grid [][]int, row int) {
	for i := 0; i < len(grid); i++ {
		if grid[i][row] == 0 {
			grid[i][row] = 1
		} else {
			grid[i][row] = 0
		}
	}
}
