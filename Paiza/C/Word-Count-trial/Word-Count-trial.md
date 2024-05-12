# 単語のカウント

<https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_00>


スペースで区切られた英単語列が与えられます。
英単語列に含まれる英単語の出現回数を出現した順番に出力してください。

> 入力例1
red green blue blue green blue

> 出力例1
red 1
green 2
blue 3

> 入力例2
Apple Apricot Orange Cherry Apple Orange Cherry Orange

> 出力例2
Apple 2
Apricot 1
Orange 3
Cherry 2

---

# Code
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	hsMap := make(map[string]int)
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		if hsMap[word] >= 1 {
			hsMap[word] += 1
		} else {
			hsMap[word] = 1
			res = append(res, word)
		}
	}

	for _, v := range res {
		v = v + " " + strconv.Itoa(hsMap[v])
		fmt.Println(v)
	}
}
```