package api

type Stack struct {
	items []int
}

func (stack *Stack) Pop() int {

	if len(stack.items) == 0 {
		return -1
	}

	elem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return elem

}

func (stack *Stack) Push(elem int) {
	stack.items = append(stack.items, elem)
}
