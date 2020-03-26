package eval

import (
	"strconv"
	"strings"

	"github.com/jms7446/mygo/tool"
)

// Eval is a simple calculotr
func Eval(expr string) int {
	nums := tool.NewStack()
	ops := tool.NewStack()
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops.Pop().(string)
			if strings.Index(higher, op) < 0 {
				ops.Push(op)
				return
			} else if op == "(" {
				return
			}
			b, a := nums.Pop().(int), nums.Pop().(int)
			var res int
			switch op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			nums = append(nums, res)
		}

	}
	for _, token := range strings.Fields(expr) {
		switch token {
		case "(":
			ops.Push(token)
		case "+", "-":
			reduce("+-*/")
			ops.Push(token)
		case "*", "/":
			reduce("*/")
			ops.Push(token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums.Push(num)
		}
	}
	reduce("+-*/")
	return nums.Pop().(int)

}
