package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func mainError() (err error) {
	h := 11.8
	l := 1.2
	r := h - l
	twelfth := r / 12

	fmt.Printf("range: %0.2f\n", r)

	clock := []float64{1, 2, 3, 3, 2, 1}
	levels := make([]float64, 0, len(clock))
	{
		level := float64(0)
		for _, c := range clock {
			level += c
			levels = append(levels, level)
		}
	}

	fmt.Println("\nleft")
	for _, level := range levels {
		fmt.Printf("%0.2f\n", twelfth*level)
	}

	fmt.Println("\nright")
	fmt.Println(h)
	fmt.Println("----")
	for _, level := range levels {
		fmt.Printf("%0.2f\n", h-twelfth*level)
	}

	return nil
}

func mainCode() int {
	err := mainError()
	if err == nil {
		return 0
	}
	fmt.Fprintf(os.Stderr, "%v: Error: %v\n", filepath.Base(os.Args[0]), err)
	return 1
}

func main() {
	os.Exit(mainCode())
}
