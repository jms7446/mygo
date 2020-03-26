package io

import (
	"bufio"
	"fmt"
	"io"
)

// WriteTo :
func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

// ReadFrom :
func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// ReadFrom2 :
func ReadFrom2(r io.Reader, f func(string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// NewIntGenerator :
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}
