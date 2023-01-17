package w20230107

type DataStream struct {
	k     int
	value int
	cnt   map[int]int
	queue []int
}

func Constructor(value int, k int) DataStream {
	d := DataStream{
		k:     k,
		value: value,
		cnt:   make(map[int]int),
		queue: make([]int, 0, k),
	}
	return d
}

func (this *DataStream) Consec(num int) bool {
	if this.k == len(this.queue) {
		this.cnt[this.queue[0]]--
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, num)
	this.cnt[num]++
	return this.cnt[this.value] == this.k
}
