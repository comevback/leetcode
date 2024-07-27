package main

import "fmt"

func main() {
	ratings := []int{1, 0, 2}
	res := candy(ratings)
	fmt.Println(res)
}

func candy(ratings []int) int {
	var Childs []Child // 创建一个Child结构体的切片，用于存储每个孩子的评分和糖果数

	// 遍历评分数组，初始化每个孩子的糖果数
	for i, v := range ratings {
		candyNum := 1 // 每个孩子至少应该有一颗糖果
		// 如果当前孩子的评分高于前一个孩子，那么他应该比前一个孩子多一颗糖果
		if i > 0 && v > Childs[i-1].rating {
			candyNum = Childs[i-1].candy + 1
		}
		// 将当前孩子的评分和糖果数添加到Childs切片中
		Childs = append(Childs, Child{rating: v, candy: candyNum})
	}

	sum := 0                           // 用于存储总糖果数
	sum += Childs[len(Childs)-1].candy // 加上最后一个孩子的糖果数

	// 从后向前遍历Childs数组，以确保每个孩子相对于他右边的孩子也得到正确的糖果数
	for i := len(Childs) - 2; i >= 0; i-- {
		// 如果当前孩子的评分高于右边的孩子，但糖果数不多于右边的孩子，需要调整
		if Childs[i].rating > Childs[i+1].rating && Childs[i].candy <= Childs[i+1].candy {
			Childs[i].candy = Childs[i+1].candy + 1
		}
		sum += Childs[i].candy // 累加糖果数
	}

	return sum // 返回总糖果数
}

// Child 结构体用于存储每个孩子的评分和糖果数
type Child struct {
	rating int
	candy  int
}
