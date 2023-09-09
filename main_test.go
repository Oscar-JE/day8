package main

import (
	"day8/matrixp"
	"testing"
)

func TestSanityCheck(t *testing.T) {
	var matrix matrixp.Matrix = initMatrix("input_short.txt")
	markSeenTrees(&matrix)
}

func TestCalcSeenTrees(t *testing.T) {
	sum := calcSeenTrees("input_short.txt")
	if sum != 21 {
		t.Errorf("wrong number of trees counted in short examples was %d should be %d", sum, 21)
	}
}

func TestViewDistance(t *testing.T) {
	var matrix matrixp.Matrix = initMatrix("input_short.txt")
	direction := matrixp.InitDown(1, 2, &matrix)
	distance := ViewDistance(direction)
	if distance != 2 {
		t.Errorf("Wrong view distance was %d shoud be %d", distance, 2)
	}
}

func TestViewScore(t *testing.T) {
	m := initMatrix("input_short.txt")
	res := ViewScore(1, 2, &m)
	if res != 4 {
		t.Errorf("wrong score : was %d should be %d", res, 4)
	}
}
