package main

import (
	"day8/matrixp"
	"day8/parser"
	"fmt"
)

func main() {
	var matrix matrixp.Matrix = initMatrix("input.txt")
	res := MaximumViewScore(&matrix)
	fmt.Println(res)
}

func initMatrix(fileName string) matrixp.Matrix {
	m := parser.ParsePositiveNumbers(fileName)
	var matrix matrixp.Matrix = matrixp.InitFromBox(m)
	return matrix
}

func calcSeenTrees(fileName string) int {
	var matrix matrixp.Matrix = initMatrix(fileName)
	seen := markSeenTrees(&matrix)
	return seen.SumElements()
}

func markSeenTrees(matrix *matrixp.Matrix) matrixp.Matrix {
	seenMatrix := matrixp.Init(matrix.GetNrRows(), matrix.GetNrCols())
	row := 0
	for row < matrix.GetNrRows() {
		fromLeft := matrixp.InitLineForward(row, matrix)
		fromRight := matrixp.InitLineBackward(row, matrix)
		seenFromPerspective(fromLeft, &seenMatrix)
		seenFromPerspective(fromRight, &seenMatrix)
		row += 1
	}
	col := 0
	for col < matrix.GetNrCols() {
		fromTop := matrixp.InitColumnForeward(col, matrix)
		fromBot := matrixp.InitColumnBackward(col, matrix)
		seenFromPerspective(fromTop, &seenMatrix)
		seenFromPerspective(fromBot, &seenMatrix)
		col += 1
	}
	return seenMatrix
}

func seenFromPerspective(iterator matrixp.Directional, seen *matrixp.Matrix) {
	maxValue := -1
	for iterator.HasNext() {
		row, col, loopValue := iterator.GetNext()
		if loopValue > maxValue {
			maxValue = loopValue
			seen.Set(row, col, 1)
		}
	}
}
