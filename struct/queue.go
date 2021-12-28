package stuct

// 循环队列
type MyCircularQueue struct {
	Size int
	data []int
	Head int
	Tail int
}

/*
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Size: k,
		data: []int{},
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsEmpty() {
		this.Head++
		this.Tail++
		this.data[this.Tail] = value
		return true
	} else if this.IsFull() {
		return false
	} else {
		tail := (this.Tail + 1) % this.Size
		this.data[tail] = value
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.data[this.Tail] = 0

}

func (this *MyCircularQueue) Front() int {

}

func (this *MyCircularQueue) Rear() int {

}

func (this *MyCircularQueue) IsEmpty() bool {

}

func (this *MyCircularQueue) IsFull() bool {

} */

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
