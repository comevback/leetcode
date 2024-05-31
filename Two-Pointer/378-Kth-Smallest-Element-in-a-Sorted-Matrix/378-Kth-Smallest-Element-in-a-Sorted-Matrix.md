# 378. Kth Smallest Element in a Sorted Matrix
Solved
Medium
Topics
Companies
Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

You must find a solution with a memory complexity better than O(n2).


Example 1:
> Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

Example 2:
> Input: matrix = [[-5]], k = 1
Output: -5
 

Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 300
-109 <= matrix[i][j] <= 109
All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
1 <= k <= n2
 

Follow up:

Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview but you may find reading this paper fun.

---

# Code
```go
package main

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	pq := NewPQ()

	// 将矩阵中的每一行都放入优先队列
	for _, v := range matrix {
		pq.push(v)
	}

	// 弹出k次，最后一次弹出的就是第k小的数
	res := []int{}
	for i := 0; i < k; i++ {
		poped := pq.pop()
		res = append(res, poped[0])
		poped = poped[1:] // 弹出的数组去掉第一个元素，然后重新放入优先队列
		if len(poped) != 0 {
			pq.push(poped)
		}
	}

	return res[len(res)-1]
}

// -----------------------------------------------------------------------------------------------------------------------
// 优先队列方法

type PQ [][]int

func NewPQ() *PQ {
	return &PQ{}
}

// 上浮方法
func (pq *PQ) swim(index int) {
	parent := index / 2

	for parent >= 1 && (*pq)[parent][0] > (*pq)[index][0] {
		(*pq)[parent], (*pq)[index] = (*pq)[index], (*pq)[parent]
		index = parent
		parent = index / 2
	}
}

// 下沉方法
func (pq *PQ) sink(index int) {
	for {
		left := 2 * index
		right := 2*index + 1
		min := index

		if left < len(*pq) && (*pq)[left][0] < (*pq)[min][0] {
			min = left
		}

		if right < len(*pq) && (*pq)[right][0] < (*pq)[min][0] {
			min = right
		}

		if min == index {
			break
		}

		(*pq)[min], (*pq)[index] = (*pq)[index], (*pq)[min]
		index = min
	}
}

// 插入方法
func (pq *PQ) push(arr []int) {
	if len(*pq) == 0 {
		(*pq) = append((*pq), []int{})
	}
	(*pq) = append((*pq), arr)
	pq.swim(len(*pq) - 1)
}

// 弹出方法
func (pq *PQ) pop() []int {

	if len(*pq) <= 1 {
		return nil
	}

	poped := (*pq)[1]
	(*pq)[1] = (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	pq.sink(1)
	return poped
}

// -----------------------------------------------------------------------------------------------------------------------
```