# 95. Unique Binary Search Trees II
Solved
Medium
Topics
Companies
Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

Example 1:

Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
Example 2:

Input: n = 1
Output: [[1]]
 

Constraints:

1 <= n <= 8

---

# 生成所有可能的唯一二叉搜索树（BST）

本文解释了如何使用 Go 语言生成由 1 到 n 组成的所有可能的唯一二叉搜索树（BST）。

## 代码解析

### 函数：generateTrees

这个函数是生成所有唯一 BST 的主函数。

```go
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return build(1, n)
}
```

1. **基础情况检查**：如果 n 为 0，返回一个空的树节点切片。
   - 这是因为当 n 为 0 时，没有数字可以用来构造 BST。
2. **递归调用 build 函数**：用于生成从 1 到 n 的所有 BST。

### 函数：build

这个函数用于生成闭区间 [lo, hi] 内所有可能的 BST。

```go
func build(lo int, hi int) []*TreeNode {
    res := []*TreeNode{}
    if lo > hi {
        res = append(res, nil)
        return res
    }

    for i := lo; i <= hi; i++ {
        leftTree := build(lo, i - 1)
        rightTree := build(i + 1, hi)
        for _, left := range leftTree {
            for _, right := range rightTree {
                root := &TreeNode{Val: i}
                root.Left = left
                root.Right = right
                res = append(res, root)
            }
        }
    }
    return res
}
```

#### 详细步骤

1. **初始化结果切片**：`res` 用于存储当前区间 [lo, hi] 内所有可能的 BST。

2. **基础情况**：如果 `lo` 大于 `hi`，返回包含一个 `nil` 元素的切片。
   - 当生成叶子节点时会进入这个条件。因为叶子节点的左右子树为空，需返回包含 `nil` 的切片以便正确生成叶子节点。

3. **遍历所有可能的根节点**：使用 `for` 循环遍历从 `lo` 到 `hi` 的所有可能值作为根节点 `i`。

4. **递归构建左右子树**：
   - `leftTree` 保存所有可能的左子树，区间为 [lo, i-1]。
   - `rightTree` 保存所有可能的右子树，区间为 [i+1, hi]。

5. **组合左右子树和根节点**：
   - 嵌套循环遍历 `leftTree` 和 `rightTree`，组合所有可能的左右子树。
   - 创建新的根节点 `root`，值为 `i`，将左子树和右子树连接到根节点。
   - 将构造好的树添加到结果切片 `res` 中。

6. **返回结果**：返回区间 [lo, hi] 内所有可能的 BST。

通过上述步骤，这段代码能够生成从 1 到 n 所有可能的唯一 BST。这样就可以利用递归和组合生成所有可能的树结构。
