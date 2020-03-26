package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println("File does not exist")
		return
	}
	defer f.Close()
	var s string
	if _, err := fmt.Fscanf(f, "%s", &s); err == nil {
		fmt.Println(s)
	}
}
