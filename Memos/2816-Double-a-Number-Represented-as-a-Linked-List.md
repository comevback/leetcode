# 2816. Double a Number Represented as a Linked List（链表翻转）（双指针）

Medium

Hint
You are given the head of a non-empty linked list representing a non-negative integer without leading zeroes.

Return the head of the linked list after doubling it.

Example 1:
![2816-1](https://assets.leetcode.com/uploads/2023/05/28/example.png)
> Input: head = [1,8,9]
Output: [3,7,8]
Explanation: The figure above corresponds to the given linked list which represents the number 189. Hence, the returned linked list represents the number 189 * 2 = 378.

Example 2:
![2816-2](https://assets.leetcode.com/uploads/2023/05/28/example2.png)
> Input: head = [9,9,9]
Output: [1,9,9,8]
Explanation: The figure above corresponds to the given linked list which represents the number 999. Hence, the returned linked list reprersents the number 999 * 2 = 1998. 
 

Constraints:
> The number of nodes in the list is in the range [1, 104]
0 <= Node.val <= 9
The input is generated such that the list represents a number that does not have leading zeros, except the number 0 itself.

---

# 思路
## 1. 链表反转
将链表反转，处理进位时，可以直接加前一位的进位值
处理完后，再将链表反转回来，返回头节点

## 2. 加虚拟头节点，两次遍历
第一次遍历，把每一位的进位情况储存到一个数组中
第二次遍历，根据进位情况，将每一位的值加上进位值，处理进位情况
最后根据虚拟头节点Val是否为0，判断是否需要删除虚拟头节点

## 3. 双指针提前处理
首先根据头节点的Val是否大于等于5，判断是否需要进位
然后在处理每一个节点时，考虑下一个节点的Val是否大于等于5，如果是，则需要进位
最后一个节点不需要处理进位，直接乘2

# 代码
```go
package main

import "fmt"

func main() {
	var head *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	res := doubleIt1(head)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 翻转链表，按位计算
// 首先设定翻转链表方法，得到翻转后的头节点
// 遍历数组，每一个节点的值乘以2，加上进位，如果大于等于10，减去10，进位加1，否则进位为0
// 如果遍历到最后一个节点，且进位为1，需要再加一个节点，值为1
// 最后将处理过的链表翻转回来，返回
func doubleIt(head *ListNode) *ListNode {

	// 翻转链表
	rseHead := reverseLink(head)
	// 保存翻转后的头节点
	saveHead := rseHead
	// 定义进位变量
	addOne := 0

	// 遍历翻转后的链表
	for rseHead != nil {
		// 计算乘2后加上进位的值，就是新的Val
		newVal := rseHead.Val*2 + addOne
		// 如果大于等于10，减去10，进位加1，否则下一次进位为0
		if newVal >= 10 {
			newVal = newVal - 10
			addOne = 1
		} else {
			addOne = 0
		}

		// 新的Val赋值给当前节点
		rseHead.Val = newVal

		// 如果遍历到最后一个节点，且进位为1，需要再加一个节点，值为1
		if rseHead.Next == nil {
			if addOne == 1 {
				var newHead *ListNode = &ListNode{
					Val: 1,
				}
				rseHead.Next = newHead
				rseHead = newHead
			} else {
				// 如果进位为0，直接break
				break
			}
		}

		// 遍历下一个节点
		rseHead = rseHead.Next
	}

	// 翻转处理后的链表，返回
	resHead := reverseLink(saveHead)
	return resHead
}

// 翻转链表方法 如有不清楚，见 (206. reverse-linked-list)
func reverseLink(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	return newHead
}

// =====================================================================================================================
// 2.从前往后遍历，保存进位数组，二次遍历加上进位
func doubleIt1(head *ListNode) *ListNode {
	// 定义虚拟头节点，虚拟节点的Next指向head
	var virtualHead *ListNode = &ListNode{Val: 0, Next: head}
	// 保存虚拟头节点，做返回用
	resHead := virtualHead
	// 保存进位数组
	addOne := []int{}
	// 每一次循环的进位值
	add := 0

	// 遍历原链表
	for head != nil {
		// 计算新的Val，如果乘2后大于等于10，减去10，进位加1，否则进位为0
		Value := head.Val * 2
		if Value >= 10 {
			Value = Value - 10
			add = 1
		} else {
			add = 0
		}

		// 保存进位到进位数组
		addOne = append(addOne, add)
		// 新的Val赋值给当前节点
		head.Val = Value
		// 遍历下一个节点
		head = head.Next
	}

	// 对于进位数组的每一个值，从虚拟头节点开始，加上进位
	for _, v := range addOne {
		virtualHead.Val += v
		virtualHead = virtualHead.Next
	}

	// 如果虚拟头节点的Val为0，去掉虚拟头节点
	if resHead.Val == 0 {
		resHead = resHead.Next
	}

	// 返回
	return resHead
}

// =====================================================================================================================
// 3. 双指针方法 from other in leetcode
// 从头节点开始，如果头节点的值大于等于5，需要加一个节点，值为1，头节点的值为头节点的值减去5
// 从头节点开始，遍历链表，如果下一个节点的值小于5，当前节点的值乘2，否则当前节点的值乘2加1
// 最后一个节点的值乘2即可
func doubleIt2(head *ListNode) *ListNode {

	ans := head
	if head.Val >= 5 {
		head = &ListNode{Val: 1, Next: head}
		ans = head
		head = head.Next
	}

	for head != nil && head.Next != nil {

		if head.Next.Val < 5 {
			head.Val = (head.Val * 2) % 10
		} else {
			head.Val = (head.Val*2 + 1) % 10
		}

		head = head.Next
	}

	head.Val = (head.Val * 2) % 10
	return ans
}
```