# Search-History 検索履歴

下記の問題をプログラミングしてみよう！
あなたが利用しているブラウザでは検索ワードの履歴を見ることができません。あなたは検索ワードの履歴を見られないのは不便だと思ったので、検索ワードの履歴を見る機能を自分でつくることにしました。

検索ワードの履歴とは次のようにつくられます。

検索ワード W が以前に入力されたことがある場合：
履歴中の W を削除する。
履歴の先頭に W を追加する。
検索ワード W が以前に入力されたことがない場合：
履歴の先頭に W を追加する。

検索ワード W が N 個与えられるので、N 個の検索ワードが与えられた後の履歴を表示するプログラムを書いてください。

# Input

入力例1
```
5
book
candy
apple
book
candy
```

出力例1
```
candy
book
apple
```

入力例2
```
6
apple
book
information
note
pen
pineapple
```

出力例2
```
pineapple
pen
note
information
book
apple
```

---

# 思路

## 1. Stack

代码：
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 读取输入，输入数字为搜索的次数
	var num int
	fmt.Scan(&num)
	// 创建一个栈
	stack := NewStack()

	// 创建一个扫描器，持续读取num次输入
	scanner := bufio.NewScanner(os.Stdin)
	// 读取输入
	for scanner.Scan() {
		// 每一行输入作为字符串
		str := scanner.Text()
		// 查找栈中是否有该字符串
		f := stack.find(str)
		// 如果有，删除该字符串
		if f != nil {
			stack.delete(f)
		}

		// 将该字符串压入栈
		stack.push(str)
		// 次数减一，如果次数小于1，退出循环
		num -= 1
		if num < 1 {
			break
		}
	}

	// 打印栈
	stack.print()
}

// 定义链表节点
type ListNode struct {
	Val  string
	Next *ListNode
}

// 定义栈
type Stack struct {
	head *ListNode
}

// 栈的构造函数
func NewStack() *Stack {
	return &Stack{}
}

// 栈的push方法
func (stack *Stack) push(value string) {
	// 创建新节点
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}
	// 新节点作为栈顶
	stack.head = newNode
}

// 栈的find方法，查找返回值为value的节点
func (stack *Stack) find(value string) *ListNode {
	// 创建一个指针指向栈顶
	current := stack.head

	// 遍历栈
	for current != nil {
		// 如果找到了值为value的节点，返回该节点
		if value == current.Val {
			return current
		}
		// 指针后移
		current = current.Next
	}

	// 没有找到返回nil
	return nil
}

// 栈的delete方法，删除节点
func (stack *Stack) delete(node *ListNode) {

	// 如果要删除的节点是栈顶节点，直接将栈顶指针指向下一个节点
	if node == stack.head {
		stack.head = stack.head.Next
		return
	}

	// 创建一个指针指向栈顶
	current := stack.head
	// 遍历栈
	for current.Next != nil {
		// 如果找到了要删除的节点，将该节点的前一个节点的Next指向该节点的Next
		if node == current.Next {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

// 栈的print方法，打印栈
func (stack *Stack) print() {
	current := stack.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
```

## 2. 用数组
