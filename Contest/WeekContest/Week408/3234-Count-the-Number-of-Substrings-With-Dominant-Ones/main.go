package main

func main() {

}

func numberOfSubstrings(s string) int {
	preOnes := make([]int, len(s)+1)
	preZeros := make([]int, len(s)+1)
	zeros, ones := 0, 0

	for i, v := range s {
		if v == '0' {
			zeros += 1
		} else if v == '1' {
			ones += 1
		}
		preOnes[i+1] = ones
		preZeros[i+1] = zeros
	}

	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			oneNums := preOnes[j+1] - preOnes[i]
			zeroNums := preZeros[j+1] - preZeros[i]
			if oneNums >= zeroNums*zeroNums {
				res += 1
			}
		}
	}

	return res
}
