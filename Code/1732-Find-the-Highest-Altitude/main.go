package main

func main() {

}

// 1.一次遍历
func largestAltitude(gain []int) int {
	// 设海拔为alti
	alti := 0
	// 达到的最高海拔为maxAlti
	maxAlti := 0

	// 遍历gain，累加海拔，如果当前海拔大于最高海拔，替换
	for _, v := range gain {
		alti += v
		if alti > maxAlti {
			maxAlti = alti
		}
	}

	// 返回最高海拔
	return maxAlti
}
