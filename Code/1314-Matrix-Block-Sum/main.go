package main

import "fmt"

func main() {
	// 初始化矩阵
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// 定义K值
	K := 1
	// 计算矩阵块和
	res := matrixBlockSum(mat, K)
	// 打印结果
	fmt.Println(res)
}

// 计算矩阵块和函数
func matrixBlockSum(mat [][]int, K int) [][]int {
	// 创建前缀和矩阵，行和列都多加1以便于计算
	prefix := make([][]int, len(mat)+1)
	for i := range prefix {
		prefix[i] = make([]int, len(mat[0])+1)
	}

	// 计算前缀和矩阵
	for i := 1; i < len(prefix); i++ {
		sum := 0
		for j := 1; j < len(prefix[0]); j++ {
			sum += mat[i-1][j-1] + prefix[i-1][j] - prefix[i-1][j-1]
			prefix[i][j] = sum
		}
	}

	// 通过前缀和矩阵计算结果矩阵
	res := getSumMatrix(mat, prefix, K)
	return res
}

// 根据前缀和矩阵和K值计算每个块的和
func getSumMatrix(nums [][]int, prefix [][]int, k int) [][]int {
	// 初始化结果矩阵
	res := make([][]int, len(nums))
	for i := range res {
		res[i] = make([]int, len(nums[0]))
	}

	// 遍历每个元素，计算其K块和
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			// 确定块的上下左右边界
			r1, c1, r2, c2 := i-k, j-k, i+k, j+k

			// 边界检查，防止超出矩阵范围
			if r1 < 0 {
				r1 = 0
			}
			if c1 < 0 {
				c1 = 0
			}

			if r2 > len(nums)-1 {
				r2 = len(nums) - 1
			}
			if c2 > len(nums[0])-1 {
				c2 = len(nums[0]) - 1
			}

			// 计算块和
			num := getSum(prefix, r1+1, c1+1, r2+1, c2+1)
			res[i][j] = num
		}
	}
	return res
}

// 利用前缀和矩阵计算指定块的和
func getSum(prefix [][]int, r1 int, c1 int, r2 int, c2 int) int {
	res := 0
	res += prefix[r2][c2]     // 加上右下角的值
	res -= prefix[r1-1][c2]   // 减去上边界外的部分
	res -= prefix[r2][c1-1]   // 减去左边界外的部分
	res += prefix[r1-1][c1-1] // 加上左上角被减去的部分
	return res
}
