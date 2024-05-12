package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hsMap := make(map[string]int)
	var res string
	order := 0

	scanner.Scan()

	strArr := strings.Split(scanner.Text(), " ")
	for _, v := range strArr {
		if _, exist := hsMap[v]; !exist {
			hsMap[v] = order
			order += 1
		}
	}

	scanner.Scan()
	res = scanner.Text()
	fmt.Println(hsMap[res])
}

// ============================================================================================================

func WordandEnterSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return len(data), data, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			return i + 1, data[:i], nil
		} else if data[i] == '\n' {
			return len(data), data[:i+1], nil
		}
	}

	return 0, nil, nil
}

func main1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(WordandEnterSplitter)
	hsMap := make(map[string]int)
	var res string
	order := 0

	for scanner.Scan() {
		word := scanner.Text()
		if _, exist := hsMap[word]; !exist {
			hsMap[word] = order
			order += 1
		}
		if strings.Contains(word, "\n") { // 检查是否包含换行符
			break
		}
	}

	if scanner.Scan() {
		res = scanner.Text()
		fmt.Println(hsMap[res])
	}
}
