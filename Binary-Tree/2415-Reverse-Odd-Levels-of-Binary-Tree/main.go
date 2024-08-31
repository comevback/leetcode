package main

func main() {

}

// TreeNode 定义了二叉树节点的结构。
type TreeNode struct {
	Val   int       // 当前节点的值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

// reverseOddLevels 反转二叉树中所有奇数层的节点值。
func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{}      // 使用队列来进行广度优先搜索
	newQueue := []*TreeNode{}   // 用于存储下一层的节点
	queue = append(queue, root) // 初始化队列
	index := 0                  // 用于跟踪当前层的层数

	// 当队列不为空时，循环处理每一层
	for len(queue) != 0 {
		// 遍历当前层的每个节点，将子节点添加到 newQueue
		for _, v := range queue {
			if v.Left != nil {
				newQueue = append(newQueue, v.Left)
			}
			if v.Right != nil {
				newQueue = append(newQueue, v.Right)
			}
		}
		// 如果当前层是奇数层（注意：层数从 0 开始，因此奇数层实际上是偶数索引），则反转该层的节点值
		if index%2 != 0 {
			reverse(queue)
		}
		// 更新队列为下一层的节点，清空 newQueue，增加层数索引
		queue = newQueue
		newQueue = []*TreeNode{}
		index += 1
	}

	return root // 返回修改后的根节点
}

// reverse 用于反转一个 TreeNode 切片中的节点值。
func reverse(arr []*TreeNode) {
	if len(arr) < 2 {
		return // 如果数组长度小于 2，不需要反转
	}

	head, tail := 0, len(arr)-1

	// 使用两个指针，一个从开始一个从末尾，交换这两个指针指向的节点的值
	for head < tail {
		arr[head].Val, arr[tail].Val = arr[tail].Val, arr[head].Val
		head += 1
		tail -= 1
	}
}
