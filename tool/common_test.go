package tool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if len(codeCount) != 3 {
		t.Error("CodeCount:", codeCount)
		t.Fatal("count should be 3 but:", len(codeCount))
	}
	if codeCount['가'] != 1 || codeCount['나'] != 2 || codeCount['다'] != 1 {
		t.Error("codeCount mismatch:", codeCount)
	}
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func Example_hasDupeRune() {
	fmt.Println(hasDupeRune("abcd"))
	fmt.Println(hasDupeRune("abcda"))
	// Output:
	// false
	// true
}
