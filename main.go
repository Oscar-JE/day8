package main

import (
	"day8/parser"
)

func main() {
	println("hej")
	m := parser.ParsePositiveNumbers("input_short.txt")
	for _, el := range m {
		for _, n := range el {
			print(n)
		}
		println("")
	}
}
