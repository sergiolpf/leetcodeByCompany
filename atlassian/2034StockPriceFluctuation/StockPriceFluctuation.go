package stockpricefluctuation

import (
	"container/heap"
)

/*
Complexidade
Tempo: O(N LogN)
Espaco: O(N)
*/
type Price struct {
	value        int
	minHeapIndex int
	maxHeapIndex int
}

type MinHeap []*Price

// Pop implements heap.Interface.
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].minHeapIndex = i
	h[j].minHeapIndex = j
}

func (h *MinHeap) Push(x any) {
	n := len(*h)
	itemm := x.(*Price)
	itemm.minHeapIndex = n
	*h = append(*h, itemm)
}

// -------------------------------------------------
type MaxHeap []*Price

// Pop implements heap.Interface.
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].maxHeapIndex = i
	h[j].maxHeapIndex = j
}

func (h *MaxHeap) Push(x any) {
	n := len(*h)
	itemm := x.(*Price)
	itemm.maxHeapIndex = n
	*h = append(*h, itemm)
}

type StockPrice struct {
	ListOfStockPrice  map[int]*Price
	maxTimeStamp      int
	maxTimeStampPrice int
	minHeap           MinHeap
	maxHeap           MaxHeap
}

func Constructor() StockPrice {
	return StockPrice{
		ListOfStockPrice: make(map[int]*Price),
	}

}

func (this *StockPrice) Update(timestamp int, price int) {

	if timestamp >= this.maxTimeStamp {
		this.maxTimeStamp = timestamp
		this.maxTimeStampPrice = price
	}

	if item, ok := this.ListOfStockPrice[timestamp]; ok {
		item.value = price
		heap.Fix(&this.minHeap, item.minHeapIndex)
		heap.Fix(&this.maxHeap, item.maxHeapIndex)
		return
	}
	itemm := &Price{value: price}
	heap.Push(&this.minHeap, itemm)
	heap.Push(&this.maxHeap, itemm)
	this.ListOfStockPrice[timestamp] = itemm
}

func (this *StockPrice) Current() int {
	return this.maxTimeStampPrice
}

func (this *StockPrice) Maximum() int {

	return this.maxHeap[0].value
}

func (this *StockPrice) Minimum() int {
	return this.minHeap[0].value
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
