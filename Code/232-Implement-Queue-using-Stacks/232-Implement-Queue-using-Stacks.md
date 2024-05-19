# 232. Implement Queue using Stacks（双栈队列）

Easy

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
 

Example 1:
> Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
> MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
 

Constraints:
> 1 <= x <= 9
At most 100 calls will be made to push, pop, peek, and empty.
All the calls to pop and peek are valid.
 

Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.

---

# 思路
- 双栈实现队列
在类型MyQueue定义两个队列，一个用来接收，一个用来输出
receiveST栈用来接收输入，按顺序存进数组（栈），当需要输出时，从transferST栈的尾部输出
如果transferST栈为空，把receiveST栈的元素依次弹出，压入transferST栈中，完毕后，从transferST栈尾部输出

# 代码
```go
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
```