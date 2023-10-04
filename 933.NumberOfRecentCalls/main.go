package main

func main() {

}

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	for i := 0; i < len(this.queue); i++ {
		if t-this.queue[i] > 3000 {
			this.queue = this.queue[1:]
			i--
		}
	}
	this.queue = append(this.queue, t)

	return len(this.queue)
}
