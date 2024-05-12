# 「単語のカウント」を解くために:part5

長さ N の文字列が半角スペース区切りで与えられます。
与えられた順に、改行区切りで出力してください。
同じ単語が複数与えられた場合は、1度だけ出力してください。
その後、与えられた単語の種類の数だけ1を改行区切りで出力してください。
だたし、その 1 は配列やリストなどに格納してください。
（続きの問題ではここに 1 ではなく、各単語の出現回数を出力しますが、ここでは準備として 1 とだけ表示します）

> 入力例1
red green blue blue green blue

> 出力例1
red
green
blue
1
1
1

> 入力例2
Apple Apricot Orange Cherry Apple Orange Cherry Orange

> 出力例2
Apple
Apricot
Orange
Cherry
1
1
1
1

---

# Code
```go
package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(v)
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(1)
	}
}
```