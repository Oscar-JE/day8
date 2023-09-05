package matrix

import (
	"errors"
)

type Matrix struct {
	nrRows int
	nrCols int
	values []int
}

func Init(nrRows int, nrCols int) Matrix {
	var values []int = []int{}
	i := 0
	for i < nrRows*nrCols {
		values = append(values, 0)
		i++
	}
	return Matrix{nrRows: nrRows, nrCols: nrCols, values: values}
}

func (this *Matrix) index(row int, col int) int {
	return row*this.nrCols + col
}

func (this *Matrix) within(row int, col int) bool {
	var inRows bool = 0 <= row && row < this.nrRows
	var inCols bool = 0 <= col && col < this.nrCols
	return inRows && inCols
}

func (this *Matrix) Set(row int, col int, v int) error {
	if this.within(row, col) {
		index := this.index(row, col)
		this.values[index] = v
		return nil
	}
	return errors.New("Set of coordinate outside matrix")
}

func (this *Matrix) Get(row int, col int) (int, error) {
	if this.within(row, col) {
		index := this.index(row, col)
		return this.values[index], nil
	}
	return 0, errors.New("Get of coordinate outside matrix")
}

func (this *Matrix) AppendRow(row []int) error {
	if len(row) == this.nrCols {
		this.values = append(this.values, row...)
		this.nrRows++
		return nil
	}
	return errors.New("appenden row length does not match column length")
}
