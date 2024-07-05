package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	res := isValidSerialization1(str)
	fmt.Println(res) // 输出 true
}

// ************************************************  array count check  ********************************************

// isValidSerialization 通过计数检查前序遍历的有效性
func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	// 将字符串分割成数组
	arr := strings.Split(preorder, ",")

	count := 0
	for i, v := range arr {
		if v != "#" {
			if i == 0 { // 如果是第一个节点，count 加 2
				count += 2
			} else { // 如果不是第一个节点，count 加 1
				count += 1
			}
		} else { // 如果是 "#"，count 减 1
			count -= 1
		}
		if count < 1 && i != len(arr)-1 { // 如果在遍历过程中count数小于1且不是最后一个节点，返回 false
			return false
		}
	}

	return count == 0 // 最终检查count数是否为0
}

// ************************************************  反序列化  ****************************************************

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool

// isValidSerialization1 通过反序列化检查前序遍历的有效性
func isValidSerialization1(preorder string) bool {
	res = true
	arr := strings.Split(preorder, ",")
	gene(&arr)
	if len(arr) == 0 && res {
		return true
	} else {
		return false
	}
}

// gene 生成二叉树，如果在生成过程中发现不合法的情况则设置 res 为 false
func gene(arr *[]string) *TreeNode {
	if len(*arr) == 0 {
		res = false
		return nil
	}

	num, err := strconv.Atoi((*arr)[0])
	(*arr) = (*arr)[1:]
	if err != nil {
		return nil
	}

	newNode := &TreeNode{
		Val: num,
	}

	newNode.Left = gene(arr)
	newNode.Right = gene(arr)
	return newNode
}

// ***********************************************  labuladong solution  ******************************************

// Solution2 定义 labuladong 方法的结构体
type Solution2 struct{}

// isValidSerialization 通过前序遍历反序列化检查前序遍历的有效性
func (s *Solution2) isValidSerialization(preorder string) bool {
	// 将字符串转化成列表
	nodes := strings.Split(preorder, ",")
	return s.deserialize(&nodes) && len(nodes) == 0
}

// deserialize 改造后的前序遍历反序列化函数
// 详细解析：https://labuladong.online/algo/data-structure/serialize-and-deserialize-binary-tree/
func (s *Solution2) deserialize(nodes *[]string) bool {
	if len(*nodes) == 0 {
		return false
	}

	// ***** 前序遍历位置 *****
	// 列表最左侧就是根节点
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == "#" {
		return true
	}
	// *********************

	return s.deserialize(nodes) && s.deserialize(nodes)
}
