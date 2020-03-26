package calc

import (
	"strconv"
	"strings"

	"github.com/jms7446/mygo/tool"
)

type BinOp func(int, int) int

func eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	var ops = tool.NewStack()
	var nums = tool.NewStack()
	ops.Push("(")

	reduce := func(newOp string) {
		// fmt.Println("Reduce called with", newOp, ops, nums)
		higher := prec[newOp]
		for len(ops) > 0 {
			op := ops.Pop().(string)
			// fmt.Printf(">>> %s\n", op)
			if op == "(" {
				if newOp != ")" {
					ops.Push(op)
				}
				return
			}
			if _, exists := higher[op]; !exists && newOp != ")" {
				ops.Push(op)
				return
			}
			f := opMap[op]
			b, a := nums.Pop().(int), nums.Pop().(int)
			// fmt.Printf(">>>  %d, %d with %s, %d\n", a, b, op, f(a, b))
			nums.Push(f(a, b))
		}

	}

	for _, token := range strings.Fields(expr) {
		if token == "(" {
			ops.Push(token)
		} else if token == ")" {
			reduce(token)
		} else if _, exists := prec[token]; exists {
			reduce(token)
			ops.Push(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums.Push(num)
		}
	}
	reduce(")")
	return nums.Pop().(int)
}

type StrSet map[string]struct{}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return eval(opMap, prec, expr)
	}
}
