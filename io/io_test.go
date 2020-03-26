package io

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string{
		"bill",
		"tom",
		"jane",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// bill
	// tom
	// jane
}

func ExampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

func ExampleReadFrom2() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom2(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )

}

func ExampleReadFrom2_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	err := ReadFrom2(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}

func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen1(), gen1(), gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1(), gen1(), gen1())
	// Output:
	// 1 2 3 4 5
	// 1 2 3 4 5
	// 6 7 8 9 10
}
