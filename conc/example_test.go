package conc

import "fmt"

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}))
	//Output: 23
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}, 3))
	//Output: 23
}
