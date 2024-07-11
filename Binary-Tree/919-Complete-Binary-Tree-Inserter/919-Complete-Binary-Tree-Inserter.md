# 919. Complete Binary Tree Inserter

Solved
Medium
Topics
Companies
A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.

Design an algorithm to insert a new node to a complete binary tree keeping it complete after the insertion.

Implement the CBTInserter class:

CBTInserter(TreeNode root) Initializes the data structure with the root of the complete binary tree.
int insert(int v) Inserts a TreeNode into the tree with value Node.val == val so that the tree remains complete, and returns the value of the parent of the inserted TreeNode.
TreeNode get_root() Returns the root node of the tree.

Example 1:

Input
["CBTInserter", "insert", "insert", "get_root"]
[[[1, 2]], [3], [4], []]
Output
[null, 1, 2, [1, 2, 3, 4]]

Explanation
CBTInserter cBTInserter = new CBTInserter([1, 2]);
cBTInserter.insert(3); // return 1
cBTInserter.insert(4); // return 2
cBTInserter.get_root(); // return [1, 2, 3, 4]

Constraints:

The number of nodes in the tree will be in the range [1, 1000].
0 <= Node.val <= 5000
root is a complete binary tree.
0 <= val <= 5000
At most 104 calls will be made to insert and get_root.

---

# Code

```go
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CBTInserter 是一个支持插入节点到完全二叉树的结构
type CBTInserter struct {
	headNode   *TreeNode   // 树的根节点
	InsertAble []*TreeNode // 保存可以插入新节点的节点切片
}

// Constructor 使用给定的树根初始化 CBTInserter
func Constructor(root *TreeNode) CBTInserter {
	InsertAble := []*TreeNode{} // 用于存储可以插入新节点的节点的切片
	queue := []*TreeNode{}      // 用于层序遍历的队列
	queue = append(queue, root) // 从根节点开始

	// 层序遍历以填充 InsertAble 切片
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			// 如果当前节点有左子节点，加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 如果当前节点有右子节点，加入队列
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 如果当前节点的左或右子节点为空，加入 InsertAble 切片
			if queue[i].Left == nil || queue[i].Right == nil {
				InsertAble = append(InsertAble, queue[i])
			}
		}
		// 去掉已处理的节点
		queue = queue[length:]
	}

	return CBTInserter{
		headNode:   root,
		InsertAble: InsertAble,
	}
}

// Insert 在完全二叉树中插入一个新节点，并返回其父节点的值
func (this *CBTInserter) Insert(val int) int {
	newNode := &TreeNode{Val: val}
	res := this.InsertAble[0].Val
	if this.InsertAble[0].Left == nil {
		this.InsertAble[0].Left = newNode
	} else if this.InsertAble[0].Right == nil {
		this.InsertAble[0].Right = newNode
		// 当一个节点的左右子节点都填满时，从 InsertAble 切片中移除
		this.InsertAble = this.InsertAble[1:]
	}

	// 将新节点加入 InsertAble 切片
	this.InsertAble = append(this.InsertAble, newNode)
	return res
}

// Get_root 返回树的根节点
func (this *CBTInserter) Get_root() *TreeNode {
	return this.headNode
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
```
