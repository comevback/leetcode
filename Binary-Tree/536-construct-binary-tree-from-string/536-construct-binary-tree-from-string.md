# 536. 从字符串生成二叉树

你需要用一个包括括号和整数的字符串构建一棵二叉树。

输入的字符串代表一棵二叉树。它包括整数和随后的 0 、1 或 2 对括号。整数代表根的值，一对括号内表示同样结构的子树。

若存在子结点，则从左子结点开始构建。

示例 1:

输入： s = "4(2(3)(1))(6(5))"
输出： [4,2,6,3,1,5]
示例 2:

输入： s = "4(2(3)(1))(6(5)(7))"
输出： [4,2,6,3,1,5,7]
示例 3:

输入： s = "-4(2(3)(1))(6(5)(7))"
输出： [-4,2,6,3,1,5,7]
提示：

0 <= s.length <= 3 \* 104
输入字符串中只包含 '(', ')', '-' 和 '0' ~ '9'
空树由 "" 而非"()"表示。

---

# Code

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "14(2(3)(1))(6(5))"
	root := str2tree(s)
	printTree(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将字符串转换为二叉树
func str2tree(s string) *TreeNode {
	return dfsBuild(s)
}

// 深度优先构建二叉树
func dfsBuild(s string) *TreeNode {
	if len(s) == 0 {
		return nil // 如果字符串为空，返回nil
	}

	// 找到第一个'('的位置，截取节点值
	numIndex := 0
	for numIndex < len(s) && s[numIndex] != '(' {
		numIndex++
	}
	intByte := s[0:numIndex] // 节点值部分的字符串

	// 将字符串转换为整数
	val, err := strconv.Atoi(intByte)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil // 如果转换出错，返回nil
	}

	// 创建新的树节点
	newNode := &TreeNode{Val: val}

	// 如果只有节点值，没有子树，直接返回节点
	if len(s) == 1 {
		return newNode
	}

	// 处理左子树
	bNum := 0
	index := numIndex
	if index < len(s) && s[index] == '(' {
		bNum++ // 遇到左括号，计数加1
	} else {
		return newNode // 没有子树，直接返回节点
	}

	// 找到左子树字符串的结束位置
	for index < len(s) && bNum != 0 {
		index++
		if index < len(s) && s[index] == '(' {
			bNum++ // 遇到左括号，计数加1
		} else if index < len(s) && s[index] == ')' {
			bNum-- // 遇到右括号，计数减1
		}
	}

	// 提取左子树字符串，并递归构建左子树
	leftStr := s[numIndex+1 : index]
	newNode.Left = dfsBuild(leftStr)

	// 处理右子树
	if index != len(s)-1 {
		// 提取右子树字符串，并递归构建右子树
		rightStr := s[index+2 : len(s)-1]
		newNode.Right = dfsBuild(rightStr)
	}

	return newNode // 返回构建好的树节点
}

// 打印二叉树（前序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		return // 如果节点为空，直接返回
	}
	fmt.Println(root.Val) // 打印当前节点值
	printTree(root.Left)  // 打印左子树
	printTree(root.Right) // 打印右子树
}
```
