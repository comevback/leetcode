# 【殿堂入りキャンペーン】みかんの仕分け

あなたはみかん農園を営んでいます。

収穫の季節になると、あなたはみかんを収穫し、重さごとに仕分けなければいけません。

仕分け作業は非常に時間がかかる作業なので、あなたはみかんを重さごとに仕分けてくれるプログラムを書くことにしました。

みかんはある定数 N の倍数のうち、正の整数の重さが書かれた箱に仕分けられます。

例えば N = 10 の時、10 g, 20 g, 30 g ... のように仕分けられます。

そして、そのみかんの重さが一番近い数の重さに仕分けられます。
重さが一番近い箱が複数ある場合、数が大きい方の箱に仕分けられます。

入力例 1 は以下のようになります。

![alt text](https://paiza-learning-mondai.s3.amazonaws.com/skillcheck_archive/mikan/img02.png)

24 g のみかんはより値の近い 20 g と書かれた箱に仕分けられます。

![alt text](https://paiza-learning-mondai.s3.amazonaws.com/skillcheck_archive/mikan/img03.png)

35 g のみかんは 30 g, 40 g の箱と差が同じです。この場合は数の大きい方の 40 g の箱へ仕分けてください。

![alt text](https://paiza-learning-mondai.s3.amazonaws.com/skillcheck_archive/mikan/img04.png)



0g に仕分けられることはないので、一番小さい重さに仕分けてください。

▼　下記解答欄にコードを記入してみよう

入力される値
入力は以下のフォーマットで与えられます。

> N M
w_1
w_2
...
w_M

・ 1 行目にそれぞれ仕分ける重さの区切りを表す整数、みかんの個数を表す整数 N, M がこの順で半角スペース区切りで与えられます。
・ 続く M 行のうちの i 行目 (1 ≦ i ≦ M) には、i 番目のみかんの重さを表す整数 w_i が与えられます。
・ 入力は合計で M + 1 行となり、入力値最終行の末尾に改行が １ つ入ります。


入力値最終行の末尾に改行が１つ入ります。
文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
期待する出力
仕分けるべき重さを以下の形式で出力してください。


> y_1
y_2
...
y_M

・ 期待する出力は M 行からなります。
・ i 行目 (1 ≦ i ≦ M) にはそれぞれ i 番目のみかんの仕分け先の重さを表す y_i を出力して下さい。
・ すべて整数で出力してください。
・ 出力最終行の末尾に改行を入れ、余計な文字、空行を含んではいけません。

条件
すべてのテストケースにおいて、以下の条件をみたします。

> ・ 1 ≦ N ≦ 100
・ 1 ≦ M ≦ 10
・ 1 ≦ w_i ≦ 1,000 (1 ≦ i ≦ M)

> 入力例1
10 3
24
35
3

> 出力例1
20
40
10

> 入力例2
50 3
40
90
10

> 出力例2
50
100
50

> 入力例3
3 2
9
5

> 出力例3
9
6

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
	var N, num int
	fmt.Scanf("%d %d", &N, &num)
	res := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		result := findNearestValue(N, weight)
		res = append(res, result)

		num -= 1
		if num < 1 {
			break
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

// 找最接近箱子功能
func findNearestValue(N int, weight int) int {
	if weight <= N {
		return N
	}

	min := N - weight
	if min < 0 {
		min = -min
	}

	times := 1
	for {
		temp := N*times - weight
		if temp < 0 {
			temp = -temp
		}
		if temp <= min {
			min = temp
			times += 1
		} else {
			return N * (times - 1)
		}
	}
}
```