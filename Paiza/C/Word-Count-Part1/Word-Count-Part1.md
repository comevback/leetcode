# 「「単語のカウント」を解くために:part1 


長さ N の文字列が半角スペース区切りで与えられます。
与えられた順に、改行区切りで出力してください。

> 入力例1
red green blue blue green blue

> 出力例1
red
green
blue
blue
green
blue

> 入力例2
Apple Apricot Orange Cherry Apple Orange Cherry Orange

> 出力例2
Apple
Apricot
Orange
Cherry
Apple
Orange
Cherry
Orange

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
	res := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		res = append(res, word)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
```