package main

import (
	"day8/matrix"
	"day8/parser"
	"fmt"
)

func main() {
	println("hej")
	m := parser.ParsePositiveNumbers("input_short.txt")
	/*
		for _, el := range m {
			for _, n := range el {
				print(n)
			}
			println("")
		}
	*/
	var matrix matrix.Matrix = matrix.InitFromBox(m)
	fmt.Println(matrix)
}
