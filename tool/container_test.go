package tool

import "testing"

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	if stack.Pop().(int) != 2 {
		t.Error("Stack pop must return recent element")
	}
}
