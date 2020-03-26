package tool

// Stack :
type Stack []interface{}

// NewStack :
func NewStack() Stack {
	return Stack{}
}

// Push :
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Pop :
func (s *Stack) Pop() interface{} {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}
