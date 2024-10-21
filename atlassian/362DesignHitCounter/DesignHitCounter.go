package designhitcounter

/*
complexidade
Tempo: O(1)
Espaco: O(N)
*/

const (
	FIVE_MINUES = 300
)

type HitCounter struct {
	hits       []int
	timestamps []int
}

func Constructor() HitCounter {
	return HitCounter{
		hits:       make([]int, 300),
		timestamps: make([]int, 300),
	}
}

func (this *HitCounter) Hit(timestamp int) {

	index := timestamp % FIVE_MINUES

	if this.timestamps[index] != timestamp {
		this.timestamps[index] = timestamp
		this.hits[index] = 1
	} else {
		this.hits[index]++
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	count := 0

	for index := 0; index < FIVE_MINUES; index++ {
		if timestamp-this.timestamps[index] < FIVE_MINUES {
			count += this.hits[index]
		}
	}

	return count
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
