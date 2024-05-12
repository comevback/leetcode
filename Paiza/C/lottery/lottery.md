# 宝くじ (paizaランク C 相当)

<https://paiza.jp/works/mondai/c_rank_skillcheck_archive/lottery>

下記の問題をプログラミングしてみよう！
今年もパイザ宝くじの季節がやってきました。パイザ宝くじ券には 100000 以上 199999 以下の 6 桁の番号がついています。毎年1つ当選番号 (100000 以上 199999 以下)が発表され、当たりクジの番号が以下のように決まります。

1等：当選番号と一致する番号
前後賞：当選番号の ±1 の番号（当選番号が 100000 または 199999 の場合，前後賞は一つしかありません）
2等：当選番号と下 4 桁が一致する番号（1等に該当する番号を除く）
3等：当選番号と下 3 桁が一致する番号（1等および2等に該当する番号を除く）

たとえば、当選番号が 142358 の場合、当たりの番号は以下のようになります。

1等：142358
前後賞：142357 と 142359
2等：102358, 112358, 122358, …, 192358 （全 9 個）
3等：100358, 101358, 102358, …, 199358 （全 90 個）

あなたが購入した n 枚の宝くじ券の各番号が入力されるので、それぞれの番号について、何等に当選したかを出力するプログラムを書いて下さい。

> 入力例1
142358
3
195283
167358
142359

> 出力例1
blank
third
adjacent

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
	var priceNumber, n int
	fmt.Scan(&priceNumber)
	fmt.Scan(&n)
	strPrice := strconv.Itoa(priceNumber)

	var res []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()

		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		if number == priceNumber {
			res = append(res, "first")
		} else if (number == priceNumber+1 || number == priceNumber-1) && number > 100000 && number < 199999 {
			res = append(res, "adjacent")
		} else if str[2:] == strPrice[2:] {
			res = append(res, "second")
		} else if str[3:] == strPrice[3:] {
			res = append(res, "third")
		} else {
			res = append(res, "blank")
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
```