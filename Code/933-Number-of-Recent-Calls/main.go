package main

func main() {

}

// =========================================================================================================
// 1. 用数组遍历，较慢
type RecentCounter1 struct {
	calls []int // 存储所有调用时间
}

// 构造函数，初始化RecentCounter1对象
func Constructor1() RecentCounter1 {
	return RecentCounter1{}
}

// Ping1方法，记录一个新的时间t，并返回最近3000毫秒内的请求数
func (this *RecentCounter1) Ping1(t int) int {
	this.calls = append(this.calls, t) // 将当前时间t添加到calls数组中
	nums := 0
	for i := len(this.calls) - 1; i >= 0; i-- { // 从后向前遍历calls数组
		if this.calls[i] >= t-3000 { // 如果时间在t-3000以内
			nums += 1 // 计数加1
		} else {
			break // 如果时间超出t-3000，停止遍历
		}
	}
	return nums // 返回计数
}

// =========================================================================================================
// 2. 用链表实现Queue
type RecentCounter struct {
	head *LinkedNode // 链表头节点
	tail *LinkedNode // 链表尾节点
}

// 构造函数，初始化RecentCounter对象
func Constructor() RecentCounter {
	return RecentCounter{}
}

// Ping方法，记录一个新的时间t，并返回最近3000毫秒内的请求数
func (this *RecentCounter) Ping(t int) int {
	var no int

	if this.head == nil { // 如果链表为空
		no = 1 // 第一个节点编号为1
	} else {
		no = this.tail.No + 1 // 新节点的编号为尾节点编号加1
	}

	// 创建新节点
	newNode := &LinkedNode{
		Val: t,
		No:  no,
	}

	if this.head == nil { // 如果链表为空
		this.head = newNode // 新节点为头节点
		this.tail = newNode // 新节点为尾节点
		return 1            // 返回1
	}

	this.tail.Next = newNode // 将新节点链接到链表末尾
	this.tail = newNode      // 更新尾节点为新节点

	// 移除所有不在最近3000毫秒内的节点
	for this.head.Val < t-3000 {
		this.head = this.head.Next
	}

	return this.tail.No - this.head.No + 1 // 返回最近3000毫秒内的请求数
}

// 链表节点结构
type LinkedNode struct {
	Val  int         // 时间值
	Next *LinkedNode // 指向下一个节点的指针
	No   int         // 节点编号
}

// =========================================================================================================
// leetcode最快解法，双指针
type RecentCounter2 struct {
	k   int //用于指示大于 t-3000 的临界点
	arr []int
}

func Constructor2() RecentCounter2 {
	return RecentCounter2{0, []int{}}
}

func (this *RecentCounter2) Ping2(t int) int {
	this.arr = append(this.arr, t)

	// 如果当前的k指向的arr元素小于t-3000，k右移
	for this.arr[this.k] < t-3000 {
		this.k++
	}

	// 返回总长度剪去不符合的个数
	return len(this.arr) - this.k
}

/**
 * 你的 RecentCounter 对象将被实例化并调用如下：
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
