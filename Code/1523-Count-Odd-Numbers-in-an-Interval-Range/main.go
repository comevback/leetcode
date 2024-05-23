package main

func main() {

}

func countOdds(low int, high int) int {
	var count int
	if low%2 == 1 || high%2 == 1 {
		count = (high-low)/2 + 1
	} else {
		count = (high - low) / 2
	}
	return count
}
