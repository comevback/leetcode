# 449. Serialize and Deserialize BST

Solved
Medium
Topics
Companies
Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary search tree. There is no restriction on how your serialization/deserialization algorithm should work. You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.

The encoded string should be as compact as possible.

Example 1:

Input: root = [2,1,3]
Output: [2,1,3]
Example 2:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 104].
0 <= Node.val <= 104
The input tree is guaranteed to be a binary search tree.

---

# Code

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 创建一个简单的二叉搜索树
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	// 创建一个Codec实例
	codec := Constructor()
	// 序列化树
	serialized := codec.serialize(root)
	fmt.Println(serialized)
	// 反序列化字符串为树
	deserialized := codec.deserialize(serialized)
	fmt.Println(deserialized)
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec 定义编码器和解码器结构
type Codec struct {
}

// Constructor 创建一个新的Codec实例
func Constructor() Codec {
	return Codec{}
}

// serialize 将树序列化为单个字符串
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	strVal := strconv.Itoa(root.Val)
	res := strVal

	if left != "" {
		res = res + "->" + left
	}

	if right != "" {
		res = res + "->" + right
	}

	return res
}

// deserialize 将编码的数据反序列化为树
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, "->")
	numArr := []int{}
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numArr = append(numArr, num)
	}

	return deser(numArr)
}

// deser 是一个递归函数，将整数数组转化为二叉树
func deser(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	newNode := &TreeNode{Val: arr[0]}
	index := 0
	// 找到第一个大于根节点值的元素下标
	for index < len(arr) && arr[index] <= arr[0] {
		index += 1
	}

	// 左子树由小于根节点值的元素构成
	if index >= 1 {
		newNode.Left = deser(arr[1:index])
	}

	// 右子树由大于根节点值的元素构成
	if index < len(arr) {
		newNode.Right = deser(arr[index:])
	}

	return newNode
}
```
