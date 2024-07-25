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
