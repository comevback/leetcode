# 297. Serialize and Deserialize Binary Tree

<https://labuladong.online/algo/data-structure/serialize-and-deserialize-binary-tree/>

Hard

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.


Example 1:
![serdeser](https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg)
> Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:
> Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000

---

## 1. DFS

```go
// ====================================================  前序遍历 DFS  ======================================================
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return preOrderStr(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// DFS
	arr := strings.Split(data, "->")
	return preOrderStrToTree(&arr)
}

func preOrderStr(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	// 用 strings.Builder 拼接字符串比用 + 拼接字符串效率更高
	// var builder strings.Builder
	// builder.WriteString(strconv.Itoa(root.Val))
	// builder.WriteString("->")
	// builder.WriteString(preOrderStr(root.Left))
	// builder.WriteString("->")
	// builder.WriteString(preOrderStr(root.Right))
	// res := builder.String()

	// 用字符串数组的方式拼接字符串，效率更高
	var strArr []string
	strArr = append(strArr, strconv.Itoa(root.Val))
	strArr = append(strArr, "->")
	strArr = append(strArr, preOrderStr(root.Left))
	strArr = append(strArr, "->")
	strArr = append(strArr, preOrderStr(root.Right))

	res := strings.Join(strArr, "")

	return res
}

// 反序列化简化
func preOrderStrToTree(arr *[]string) *TreeNode {
	if len(*arr) == 0 {
		return nil
	}

	num, err := strconv.Atoi((*arr)[0])
	*arr = (*arr)[1:]
	if err != nil {
		return nil
	}
	var root *TreeNode = &TreeNode{
		Val: num,
	}

	root.Left = preOrderStrToTree(arr)
	root.Right = preOrderStrToTree(arr)

	return root
}

// 我的反序列化方法，需要界定左右子树在数组中的范围，不用传入数组指针
func preOrderStrToTree1(arr []string) *TreeNode {
	if len(arr) == 0 || arr[0] == "#" {
		return nil
	}
	num, _ := strconv.Atoi(arr[0])
	arr = arr[1:]
	var root *TreeNode = &TreeNode{
		Val: num,
	}

	count := 0
	var rightStart int
	for i := 0; i < len(arr); i++ {
		if count == -1 {
			rightStart = i
			break
		}
		if count == 0 && i != 0 {
			rightStart = i
			break
		}
		if arr[i] == "#" {
			count -= 1
		} else {
			if i == 0 {
				count += 2
			} else {
				count += 1
			}

		}
	}

	root.Left = preOrderStrToTree1(arr[:rightStart])
	root.Right = preOrderStrToTree1(arr[rightStart:])

	return root
}
```

## 2. BFS

```go
// ====================================================  层序遍历 BFS  ======================================================
type Codec_BFS struct {
}

func Constructor_BFS() Codec_BFS {
	return Codec_BFS{}
}

// Serializes a tree to a single string.
func (this *Codec_BFS) serialize_BFS(root *TreeNode) string {
	return levelOrderStr(root)
}

// Deserializes your encoded data to tree.
func (this *Codec_BFS) deserialize_BFS(data string) *TreeNode {
	// BFS
	arr := strings.Split(data, "->")
	res := levelOrderStrToTree(arr)
	return res
}

// 层序遍历序列化和反序列化
func levelOrderStr(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	res := []string{}

	for len(queue) != 0 {
		temp := []*TreeNode{}
		for _, v := range queue {
			if v != nil {
				res = append(res, strconv.Itoa(v.Val))
				if v.Left != nil {
					temp = append(temp, v.Left)
				} else {
					temp = append(temp, nil)
				}
				if v.Right != nil {
					temp = append(temp, v.Right)
				} else {
					temp = append(temp, nil)
				}
			} else {
				res = append(res, "#")
			}
		}
		queue = temp
	}

	result := strings.Join(res, "->")
	return result
}

func levelOrderStrToTree(arr []string) *TreeNode {
	if len(arr) == 0 || arr[0] == "#" {
		return nil
	}

	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			if index < len(arr) && arr[index] != "#" {
				val, _ := strconv.Atoi(arr[index])
				node.Left = &TreeNode{Val: val}
				nextQueue = append(nextQueue, node.Left)
			}
			index += 1
			if index < len(arr) && arr[index] != "#" {
				val, _ := strconv.Atoi(arr[index])
				node.Right = &TreeNode{Val: val}
				nextQueue = append(nextQueue, node.Right)
			}
			index += 1
		}
		queue = nextQueue
	}

	return root
}
```