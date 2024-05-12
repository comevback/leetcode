package main

import (
	"errors"
	"fmt"

	queue "github.com/comevback/leetcode/DSA/Stack-Queue/Queue"
	stack "github.com/comevback/leetcode/DSA/Stack-Queue/Stack"
)

func main() {
	var binTree = NewBinary()
	binTree.insert(5)
	binTree.insert(6)
	binTree.insert(7)
	binTree.insert(8)
	binTree.insert(4)
	binTree.insert(2)

	binTree.insert(3)

	//preOrder(binTree.root)
	binTree.reversal()
}

type TreeNode struct {
	Val       int
	Frequency int
	Left      *TreeNode
	Right     *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinary() *BinaryTree {
	return &BinaryTree{}
}

// ====================================================  insert 方法  =================================================
func (binTree *BinaryTree) insert(value int) {
	newNode := &TreeNode{ // 创建一个新节点，Val为value，Frequency初始为1
		Val:       value,
		Frequency: 1,
	}

	if binTree.root == nil { // 如果根节点为空，新节点成为根节点
		binTree.root = newNode
		return
	}

	current := binTree.root // 设current指向根节点
	for {                   // 开始循环查找插入位置
		if value < current.Val { // 如果插入值小于当前节点的值
			if current.Left == nil { // 并且左子节点为空
				current.Left = newNode // 新节点成为左子节点
				break
			}
			current = current.Left // 否则移动到左子节点继续查找
		} else if value > current.Val { // 如果插入值大于当前节点的值
			if current.Right == nil { // 并且右子节点为空
				current.Right = newNode // 新节点成为右子节点
				break
			}
			current = current.Right // 否则移动到右子节点继续查找
		} else { // 如果插入值等于当前节点的值
			current.Frequency += 1 // 当前节点的频率加一
			break
		}
	}
}

// ====================================================  lookup 方法  =================================================
func (binTree *BinaryTree) lookup(value int) (*TreeNode, error) {
	current := binTree.root // 从根节点开始查找

	for current != nil { // 当当前节点不为空时
		if value < current.Val { // 如果查找值小于当前节点的值
			current = current.Left // 移动到左子节点
		} else if value > current.Val { // 如果查找值大于当前节点的值
			current = current.Right // 移动到右子节点
		} else { // 如果查找值等于当前节点的值
			return current, nil // 返回找到的节点和nil错误
		}
	}

	return nil, errors.New("didn't find") // 如果未找到，返回nil和错误信息
}

// ====================================================  remove 方法  =================================================
func (binTree *BinaryTree) remove(value int) error {
	// 定义当前节点current，父节点parentNode，替换者replacer，替换者replacerParent的父节点
	current := binTree.root
	var parentNode *TreeNode
	var replacer, replacerParent *TreeNode

	// 如果树为空，返回错误
	if current == nil {
		return errors.New("empty tree")
	}

	// 寻找要删除的节点，如果找到，parentNode跟踪current，为要删除节点的父节点
	for current != nil {
		if value < current.Val {
			parentNode = current
			current = current.Left
		} else if value > current.Val {
			parentNode = current
			current = current.Right
		} else {
			// 如果找到要删除的节点的频率大于1，频率减一，返回nil
			if current.Frequency > 1 {
				current.Frequency -= 1
				return nil
			}
			break
		}
	}

	// 如果找到的节点为空，说明没有这个值，返回错误
	if current == nil {
		return errors.New("didn't find the value")
	}

	// =============================================================================
	// 如果找到的节点是叶子节点，直接删除
	if current.Left == nil && current.Right == nil {
		if parentNode == nil {
			binTree.root = nil
		} else if parentNode.Val > current.Val {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	} else if current.Right == nil {
		// 如果找到的节点只有左子节点，直接删除
		if parentNode == nil {
			binTree.root = binTree.root.Left
		} else if parentNode.Val > current.Val {
			parentNode.Left = current.Left
		} else {
			parentNode.Right = current.Left
		}
	} else {
		// 如果找到的节点有两个子节点
		replacerParent = current
		replacer = current.Right

		// 寻找右子树中的最小节点作为替换者
		for replacer.Left != nil {
			replacerParent = replacer
			replacer = replacer.Left
		}

		// 如果替换者不是删除节点的直接右子节点
		if replacerParent != current {
			replacerParent.Left = replacer.Right // 将替换者父节点的左链接指向替换者的右子节点
			replacer.Right = current.Right       // 替换者右子链接指向当前节点的右子节点
		}
		replacer.Left = current.Left // 替换者左子链接指向当前节点的左子节点

		// 重新链接 parentNode到 replacer
		if parentNode == nil {
			binTree.root = replacer // 如果 parentNode 为空，即当前节点是根节点
		} else if parentNode.Left == current {
			parentNode.Left = replacer
		} else {
			parentNode.Right = replacer
		}
	}
	return nil
}

// ====================================================  深度优先遍历  =================================================

func (binTree *BinaryTree) reversal() {
	q1 := queue.NewQueue[int]()
	q2 := queue.NewQueue[int]()
	q3 := queue.NewQueue[int]()

	preTravel(binTree.root, q1)

	inTravel(binTree.root, q2)
	postTravel(binTree.root, q3)

	q1.PrintQueue()
	preOrder(binTree.root)
	q2.PrintQueue()
	q3.PrintQueue()
	postOrder(binTree.root)
}

// ====================================================  递归  ====================================================

// 前序遍历
func preTravel(treeNode *TreeNode, q *queue.Queue[int]) {
	q.Enqueue(treeNode.Val)
	if treeNode.Left != nil {
		preTravel(treeNode.Left, q)
	}

	if treeNode.Right != nil {
		preTravel(treeNode.Right, q)
	}
}

// 中序遍历
func inTravel(treeNode *TreeNode, q *queue.Queue[int]) {

	if treeNode.Left != nil {
		inTravel(treeNode.Left, q)
	}

	q.Enqueue(treeNode.Val)
	if treeNode.Right != nil {
		inTravel(treeNode.Right, q)
	}
}

// 后序遍历
func postTravel(treeNode *TreeNode, q *queue.Queue[int]) {

	if treeNode.Left != nil {
		postTravel(treeNode.Left, q)
	}

	if treeNode.Right != nil {
		postTravel(treeNode.Right, q)
	}
	q.Enqueue(treeNode.Val)
}

// ====================================================  迭代  ====================================================
// 前序遍历
func preOrder(treeNode *TreeNode) {
	// 定义一个栈，类型为*TreeNode
	st := stack.NewStack_Link[*TreeNode]()
	// 初始时向栈里推入根节点
	st.Push(treeNode)
	// 定义结果数组
	res := []int{}

	// 当栈不为空时循环
	for !st.IsEmpty() {
		// 每次弹出一个元素，把这个元素的Val加入到结果数组中
		poped, _ := st.Pop()
		res = append(res, poped.Val)

		// 首先如果这个元素有右子节点，压入栈
		if poped.Right != nil {
			st.Push(poped.Right)
		}
		// 然后如果这个元素有左子节点，压入栈
		if poped.Left != nil {
			st.Push(poped.Left)
		}
	}

	// 打印结果数组
	fmt.Println(res)
}

// 后序遍历
// 和前序遍历类似，因为前序遍历是：中，左，右，而后续遍历是：左，右，中
// 在前序遍历的基础上切换次序，把中，左，右，变成中，右，左，得到结果数组后，翻转结果数组，顺序就为左，右，中
func postOrder(treeNode *TreeNode) {
	// 定义一个栈，类型为*TreeNode
	st := stack.NewStack_Link[*TreeNode]()
	// 初始时向栈里推入根节点
	st.Push(treeNode)
	// 定义结果数组
	res := []int{}

	// 当栈不为空时循环
	for !st.IsEmpty() {
		// 每次弹出一个元素，把这个元素的Val加入到结果数组中
		poped, _ := st.Pop()
		res = append(res, poped.Val)

		// 首先如果这个元素有左子节点，压入栈
		if poped.Left != nil {
			st.Push(poped.Left)
		}
		// 然后如果这个元素有右子节点，压入栈
		if poped.Right != nil {
			st.Push(poped.Right)
		}
	}

	// 翻转结果数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	fmt.Println(res)
}

// 中序遍历
func inOrder(treeNode *TreeNode) {
	st := stack.NewStack_Link[*TreeNode]()
	st.Push(treeNode)
	res := []int{}

	current := treeNode

	for current.Left != nil {
		st.Push(current)
	}

	for !st.IsEmpty() {

	}

	fmt.Println(res)
}
