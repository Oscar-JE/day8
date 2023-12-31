package main

import (
	"day8/matrixp"
)

func ViewDistance(direction matrixp.Directional) int {
	dist := 0
	startValue := -1
	if direction.HasNext() {
		_, _, v := direction.GetNext()
		startValue = v
	}
	for direction.HasNext() {
		_, _, v := direction.GetNext()
		dist++
		if startValue <= v {
			break
		}
	}
	return dist
}

func ViewScore(row int, col int, matrix *matrixp.Matrix) int {
	up := matrixp.InitUp(row, col, matrix)
	right := matrixp.InitRight(row, col, matrix)
	down := matrixp.InitDown(row, col, matrix)
	left := matrixp.InitLeft(row, col, matrix)
	return ViewDistance(up) * ViewDistance(right) *
		ViewDistance(down) * ViewDistance(left)
}

func MaximumViewScore(matrix *matrixp.Matrix) int {
	row := 0
	maximum := -1
	for row < matrix.GetNrRows() {
		col := 0
		for col < matrix.GetNrCols() {
			score := ViewScore(row, col, matrix)
			maximum = max(maximum, score)
			col++
		}
		row++
	}
	return maximum
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
