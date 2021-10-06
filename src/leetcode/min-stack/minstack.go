package min_stack

type MinStack struct {
	stack1   []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack1:   []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack1 = append(this.stack1, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else if this.minStack[len(this.minStack)-1] < val {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack1 = this.stack1[:len(this.stack1)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
