package graph

import (
	"fmt"
	"io"
)

// WriteTo :
func WriteTo(w io.Writer, adjList [][]int) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for _, edges := range adjList {
		eSize := len(edges)
		if _, err := fmt.Fprintf(w, "\n%d", eSize); err != nil {
			return err
		}
		for _, e := range edges {
			if _, err := fmt.Fprintf(w, " %d", e); err != nil {
				return err
			}
		}

	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}

	return nil
}

// ReadFrom :
func ReadFrom(r io.Reader, adjList *[][]int) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}
	*adjList = make([][]int, size)
	for i := 0; i < size; i++ {
		var lsize int
		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize)
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}
	return nil
}
