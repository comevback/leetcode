package main

func main() {

}

type LFUCache struct {
	hsMap map[int]*Item
	arr   []Item
}

type Item struct {
	Key  int
	Val  int
	Freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		hsMap: make(map[int]*Item),
		arr:   []Item{},
	}
}

func (this *LFUCache) Get(key int) int {
	if _, exist := this.hsMap[key]; exist {
		this.hsMap[key].Freq += 1
		return this.hsMap[key].Val
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {

}

func (this *LFUCache) ReSort() {

}
