package main

import "fmt"

func main() {
	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	key := len(arr) - 1 // 这里是选择最后一个元素作为轴，如果用的是其他的元素，可以把这个元素和最后一个元素调换，然后开始遍历
	head := 0
	for head < key {
		if arr[head] > arr[key] {
			if head+1 == key {
				arr[key], arr[head] = arr[head], arr[key]
			} else {
				arr[key], arr[head], arr[key-1] = arr[head], arr[key-1], arr[key]
			}
			key -= 1
		} else {
			head += 1
		}
	}

	if key > 0 {
		QuickSort(arr[:key])
	}
	if key < len(arr)-1 {
		QuickSort(arr[key+1:])
	}
}
