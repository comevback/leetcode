# ソート

N 人の人々がおり、それぞれの人は金と銀を何キログラムか持っています。今は金の方が銀よりも価値が高いですが、ある日金と銀の価値が逆転して、人々の財産の多さは次のように決定されるようになりました。

1. 持っている銀が多い方が財産が多い。
2. 持っている銀が同じなら、持っている金が多い方が財産が多い。

それぞれの人が持っている金と銀のキログラム数が与えられるので、この規則にしたがって、財産を多い順に並び替えて出力してください。

入力される値
入力は以下のフォーマットで与えられます。

> N
g_1 s_1
...
g_N s_N

1 行目には人々の数を表す整数 N が与えられ、 2 行目から (N + 1) 行目には、人々が持っている金の量 g_i と銀の量 s_i がそれぞれ半角スペース区切りで N 行与えられます （1 ≤ i ≤ N）。

入力値最終行の末尾に改行が１つ入ります。
文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
期待する出力
上の規則に従って人々の財産を並び替え、入力と同じ形式で、各 g_i, s_i を半角スペース区切りで、財産が多い順に N 行出力してください。
末尾に改行を入れ、余計な文字、空行を含んではいけません。

条件
すべてのテストケースにおいて、以下の条件をみたします。

・1 ≤ N ≤ 50
・0 ≤ g_i, s_i ≤ 50（1 ≤ i ≤ N）

> 入力例1
2
2 1
1 2

> 出力例1
1 2
2 1

> 入力例2
4
2 3
0 4
5 0
3 3

> 出力例2
0 4
3 3
2 3
5 0

---

# Code
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)      // 读取用户输入的行数N
	var res *ListNode // 声明链表的头节点

	scanner := bufio.NewScanner(os.Stdin) // 创建从标准输入读取的Scanner
	for scanner.Scan() {                  // 循环读取每一行输入
		str := scanner.Text() // 获取输入的字符串
		newNode := &ListNode{ // 创建一个新的链表节点
			Val: str,
		}
		thisAmount := count(str) // 计算当前字符串的权重

		current := res      // 从链表头开始遍历
		if current == nil { // 如果链表为空，则直接插入新节点
			res = newNode
			N -= 1
			if N < 1 { // 如果已达到输入的行数，终止循环
				break
			}
			continue
		}

		var parentNode *ListNode // 用于记录当前节点的前一个节点
		for current != nil {
			Amount := count(current.Val) // 计算当前遍历到的节点的权重
			// 比较权重，决定插入位置
			if thisAmount[1] > Amount[1] || thisAmount[1] == Amount[1] && thisAmount[0] > Amount[0] {
				break
			}
			parentNode = current // 向前移动
			current = current.Next
		}

		// 插入新节点
		if parentNode == nil {
			newNode.Next = res
			res = newNode
		} else {
			newNode.Next = current
			parentNode.Next = newNode
		}

		N -= 1
		if N < 1 {
			break
		}
	}

	current := res // 从头节点开始输出链表
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// 把字符串转换为一个两个元素的整数型切片，方便比较
func count(str string) []int {
	fort := strings.Split(str, " ") // 将字符串按空格分割
	res := []int{}
	for _, v := range fort {
		num, err := strconv.Atoi(v) // 将字符串转换为整数
		if err != nil {
			panic(err) // 转换错误则抛出异常
		}
		res = append(res, num) // 将转换后的整数加入数组
	}

	return res
}

// 链表定义
type ListNode struct {
	Val  string
	Next *ListNode
}
```