# 901. Online Stock Span

Solved
Medium
Topics
Companies
Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.

For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.
Implement the StockSpanner class:

StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.

Example 1:

Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]

Explanation
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80); // return 1
stockSpanner.next(60); // return 1
stockSpanner.next(70); // return 2
stockSpanner.next(60); // return 1
stockSpanner.next(75); // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85); // return 6

Constraints:

1 <= price <= 105
At most 104 calls will be made to next.

---

# Code

```go
package main

func main() {
	obj := Constructor()
	println(obj.Next(31))
	println(obj.Next(41))
	println(obj.Next(48))
	println(obj.Next(59))
	println(obj.Next(79))
}

// StockSpanner 结构包含两个数组：一个存储价格，一个存储每个价格之前第一个比它高的价格的索引。
type StockSpanner struct {
	arr       []int // 用于存储传入的价格
	preHigher []int // 存储每个价格之前第一个比它大的价格的索引
}

// Constructor 初始化并返回一个StockSpanner对象。
func Constructor() StockSpanner {
	return StockSpanner{
		arr:       []int{},
		preHigher: []int{},
	}
}

// Next 方法接收一个价格，返回从该价格往前数，连续的价格都低于或等于该价格的天数。
func (this *StockSpanner) Next(price int) int {
	this.arr = append(this.arr, price) // 将新价格添加到数组
	index := len(this.arr) - 2         // 从最近的一个价格开始向前检查

	// 向前遍历，找到第一个高于当前价格的位置
	for index >= 0 {
		if this.arr[index] <= price {
			if len(this.preHigher) <= index {
				this.preHigher = append(this.preHigher, index)
				break
			} else {
				index = this.preHigher[index]
			}
		} else {
			this.preHigher = append(this.preHigher, index)
			break
		}
	}

	// 如果没有找到比当前价格高的，记录为-1
	if index < 0 {
		this.preHigher = append(this.preHigher, -1)
	}

	return len(this.arr) - 1 - index // 返回当前价格的跨度
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
```
