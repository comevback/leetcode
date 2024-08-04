package main

func main() {

}

func canAliceWin(nums []int) bool {
	sum := 0
	singleSum := 0

	for _, v := range nums {
		if v < 10 {
			singleSum += v
		}
		sum += v
	}

	if singleSum*2 != sum {
		return true
	} else {
		return false
	}
}
