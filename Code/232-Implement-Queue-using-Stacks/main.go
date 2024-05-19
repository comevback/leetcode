package main

import "fmt"

func main() {
	var queue MyQueue = Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}

// 1. 双栈实现队列
// 在类型MyQueue定义两个队列，一个用来接收，一个用来输出
// receiveST栈用来接收输入，按顺序存进数组（栈），当需要输出时，从transferST栈的尾部输出
// 如果transferST栈为空，把receiveST栈的元素依次弹出，压入transferST栈中，完毕后，从transferST栈尾部输出
type MyQueue struct {
	receiveST  []int
	transferST []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.receiveST = append(this.receiveST, x)
}

func (this *MyQueue) Pop() int {
	if len(this.transferST) == 0 {
		if len(this.receiveST) == 0 {
			return -1
		} else {
			for len(this.receiveST) != 0 {
				this.transferST = append(this.transferST, this.receiveST[len(this.receiveST)-1])
				this.receiveST = this.receiveST[:len(this.receiveST)-1]
			}
		}
	}
	res := this.transferST[len(this.transferST)-1]
	this.transferST = this.transferST[:len(this.transferST)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.transferST) == 0 {
		if len(this.receiveST) == 0 {
			return -1
		} else {
			for len(this.receiveST) != 0 {
				this.transferST = append(this.transferST, this.receiveST[len(this.receiveST)-1])
				this.receiveST = this.receiveST[:len(this.receiveST)-1]
			}
		}
	}
	res := this.transferST[len(this.transferST)-1]
	return res
}

func (this *MyQueue) Empty() bool {
	if len(this.receiveST) == 0 && len(this.transferST) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
