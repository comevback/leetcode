package main

func main() {

}

// 双指针
func reverseString(s []byte) {
	var head, tail int = 0, len(s) - 1

	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head += 1
		tail -= 1
	}
}
