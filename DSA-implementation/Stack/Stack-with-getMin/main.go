package main

func main() {

}

type Stack[T any] struct {
	value []T
}

type Stack_getMin struct {
}

// func (stack *Stack[T]) push(newValue any) {
// 	stack.value = append(stack.value, newValue)
// }

// func (stack *Stack) pop() any {
// 	return stack.value[len(stack.value)-1]
// }

// func (stack *Stack) getMin() any {
// 	return stack.min
// }
