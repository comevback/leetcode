# 1261. Find Elements in a Contaminated Binary Tree
Solved
Medium
Topics
Companies
Hint
Given a binary tree with the following rules:

root.val == 0
If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 * x + 2
Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.

Implement the FindElements class:

FindElements(TreeNode* root) Initializes the object with a contaminated binary tree and recovers it.
bool find(int target) Returns true if the target value exists in the recovered binary tree. 

Example 1:

Input
["FindElements","find","find"]
[[[-1,null,-1]],[1],[2]]
Output
[null,false,true]
Explanation
FindElements findElements = new FindElements([-1,null,-1]); 
findElements.find(1); // return False 
findElements.find(2); // return True 
Example 2:

Input
["FindElements","find","find","find"]
[[[-1,-1,-1,-1,-1]],[1],[3],[5]]
Output
[null,true,true,false]
Explanation
FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
findElements.find(1); // return True
findElements.find(3); // return True
findElements.find(5); // return False
Example 3:

Input
["FindElements","find","find","find","find"]
[[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
Output
[null,true,false,false,true]
Explanation
FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
findElements.find(2); // return True
findElements.find(3); // return False
findElements.find(4); // return False
findElements.find(5); // return True 

Constraints:

TreeNode.val == -1
The height of the binary tree is less than or equal to 20
The total number of nodes is between [1, 104]
Total calls of find() is between [1, 104]
0 <= target <= 106

---

# Code
```go
package main

func main() {

}

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindElements 用于处理被污染的二叉树，并支持查找元素
type FindElements struct {
	tree  *TreeNode
	hsMap map[int]bool
}

// Constructor 初始化 FindElements 结构体，修复二叉树并建立值的哈希集合
func Constructor(root *TreeNode) FindElements {
	hsMap := make(map[int]bool)
	travel(root, 0, hsMap)
	return FindElements{
		tree:  root,
		hsMap: hsMap,
	}
}

// travel 递归地遍历二叉树，修复节点的值，并将值存入哈希集合
func travel(root *TreeNode, value int, hsMap map[int]bool) {
	if root == nil {
		return
	}

	// 修复当前节点的值
	root.Val = value
	// 将修复后的值存入哈希集合
	hsMap[root.Val] = true

	// 递归修复左子节点，值为当前值 * 2 + 1
	travel(root.Left, value*2+1, hsMap)
	// 递归修复右子节点，值为当前值 * 2 + 2
	travel(root.Right, value*2+2, hsMap)
}

// Find 在修复后的二叉树中查找目标值
func (this *FindElements) Find(target int) bool {
	// 在哈希集合中查找目标值
	return this.hsMap[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
```