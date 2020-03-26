package eval

import (
	"testing"
)

func TestEval(t *testing.T) {
	var ret int

	ret = Eval("5")
	if ret != 5 {
		t.Error("base", ret)
	}
	ret = Eval("1 + 2")
	if ret != 3 {
		t.Error("simple sum", ret)
	}
	ret = Eval("1 - 2 + 3")
	if ret != 2 {
		t.Error("plus and minus", ret)
	}
	ret = Eval("3 * ( 3 + 1 * 3 ) / 2")
	if ret != 9 {
		t.Error("include par", ret)
	}
	ret = Eval("3 * ( ( 3 + 1 ) * 3 ) / 2")
	if ret != 18 {
		t.Error("include double par", ret)
	}
}
