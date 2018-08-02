package stack

var a Stack

type Stack struct {
	mass []int
}

func New() *Stack {
	a.mass = make([]int, 0)
	return &a
}

func (a *Stack) Push(i int) {
	a.mass = append(a.mass, i)
}

func (a *Stack) Pop() int {
	x := a.mass[len(a.mass)-1]
	a.mass = a.mass[:len(a.mass)-1]
	return x
}
