# 「単語のカウント」を解くために:part2

長さ N の文字列が半角スペース区切りで与えられます。
red が登場するかどうか調べてみましょう。
redが登場すれば Yes を、登場しなければ No を出力してください。

> 入力例1
red green blue blue green blue

> 出力例1
Yes

> 入力例2
Apple Apricot Orange Cherry Apple Orange Cherry Orange

> 出力例2
No

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
	res := "No"

	for scanner.Scan() {
		word := scanner.Text()
		if word == "red" {
			res = "Yes"
		}
	}

	fmt.Println(res)
}
```