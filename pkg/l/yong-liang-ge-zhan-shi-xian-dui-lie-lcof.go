/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 22:37:10
 */
package l

type CQueue struct {
	stack1 stack2
	stack2 stack2
}

func Constructor2() CQueue {
	return CQueue{
		stack1: NewStack2(),
		stack2: NewStack2(),
	}
}

func (this *CQueue) AppendTail(value int) {
	for this.stack2.len() != 0 {
		node := this.stack2.pop()
		this.stack1.push(node)
	}
	this.stack1.push(value)
}

func (this *CQueue) DeleteHead() int {
	for this.stack1.len() != 0 {
		node := this.stack1.pop()
		this.stack2.push(node)
	}
	return this.stack2.pop()
}

type stack2 struct {
	data []int
}

func NewStack2() stack2 {
	return stack2{
		data: []int{},
	}
}

func (this *stack2) len() int {
	return len(this.data)
}

func (this *stack2) push(val int) {
	this.data = append(this.data, val)
}

func (this *stack2) pop() int {
	if len(this.data) == 0 {
		return -1
	}
	res := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
