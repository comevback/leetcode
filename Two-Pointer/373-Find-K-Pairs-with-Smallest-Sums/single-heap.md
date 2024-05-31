```go
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := NewPQ()

	for _, v := range nums1 {
		for _, w := range nums2 {
			newItem := Item{
				Arr: []int{v, w},
				Val: v + w,
			}
			pq.push(newItem)
		}
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		poped := pq.pop()
		res = append(res, poped.Arr)
	}

	return res
}

type Item struct {
	Arr []int
	Val int
}

type PQ []Item

func NewPQ() *PQ {
	return &PQ{}
}

func (pq *PQ) swim(index int) {
	parent := index / 2

	for parent >= 1 && (*pq)[parent].Val > (*pq)[index].Val {
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

		for left < len(*pq) && (*pq)[left].Val < (*pq)[min].Val {
			min = left
		}

		for right < len(*pq) && (*pq)[right].Val < (*pq)[min].Val {

			min = right
		}

		if min == index {
			break
		}

		(*pq)[index], (*pq)[min] = (*pq)[min], (*pq)[index]
		index = min
	}
}

func (pq *PQ) push(item Item) {
	if len(*pq) == 0 {
		(*pq) = append((*pq), Item{})
	}

	(*pq) = append((*pq), item)
	pq.swim(len(*pq) - 1)
}

func (pq *PQ) pop() Item {
	if len(*pq) <= 1 {
		var null Item
		return null
	}

	poped := (*pq)[1]
	(*pq)[1] = (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	pq.sink(1)
	return poped
}
```