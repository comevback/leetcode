package main

import "fmt"

func main() {

	var arr1 = make([]int, 5)
	arr1 = []int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	var arr2 = New(1, 4, 5, 6, 7)
	arr2.push(9)
	fmt.Println(arr2.find(5))
	fmt.Println(arr2.print())
	arr2.delete(3)
	fmt.Println(arr2.print())
}

type MyArray struct {
	body   []int
	length int
}

func New(nums ...int) *MyArray {
	return &MyArray{
		body:   nums,
		length: len(nums),
	}
}

func (arr *MyArray) push(num int) {
	arr.body = append(arr.body, num)
	arr.length += 1
}

func (arr *MyArray) find(idx int) int {
	return arr.body[idx]
}

func (arr *MyArray) print() []int {
	return arr.body
}

func (arr *MyArray) delete(idx int) {
	arr.body = append(arr.body[:idx], arr.body[idx+1:]...)
	arr.length -= 1
}
