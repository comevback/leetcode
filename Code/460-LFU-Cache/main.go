package main

func main() {

}

type LFUCache struct {
	hsMap map[int]*Item
	arr   []Item
}

type Item struct {
	Key int
	Val int
	Fre int
}

func Constructor(capacity int) LFUCache {

}

func (this *LFUCache) Get(key int) int {

}

func (this *LFUCache) Put(key int, value int) {

}
