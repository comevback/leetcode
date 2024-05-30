package main

import "math"

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	pq := NewPQ()

	for _, v := range matrix {
		pq.push(v)
	}

	res := []int{}
	for i := 0; i < k; i++ {
		poped := pq.pop()
		res = append(res, poped[0])
		poped = poped[1:]
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

func (pq *PQ) swim(index int) {
	parent := index / 2

	for parent >= 1 && (*pq)[parent][0] > (*pq)[index][0] {
		(*pq)[parent], (*pq)[index] = (*pq)[index], (*pq)[parent]
		index = parent
		parent = index / 2
	}
}

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

func (pq *PQ) push(arr []int) {
	if len(*pq) == 0 {
		(*pq) = append((*pq), []int{})
	}
	(*pq) = append((*pq), arr)
	pq.swim(len(*pq) - 1)
}

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
func kthSmallest1(matrix [][]int, k int) int {
	length := len(matrix[0])
	arrY := make([]int, length)

	// p1, p2, p3 := []int{0, y1}, []int{0, y2}, []int{0, y3}

}

func getMinandIndex(arr []int) (int, int) {
	min := math.MaxInt
	index := -1

	for i, v := range arr {
		min = v
		index = i
	}

	return min, index
}
