package matrixp

import (
	"errors"
	"strconv"
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

func InitFromBox(numberBox [][]int) Matrix {
	nrCols := len(numberBox[0][:])
	nrLines := len(numberBox)
	m := Init(nrLines, nrCols)
	for i, line := range numberBox {
		for j, el := range line {
			m.Set(i, j, el)
		}
	}
	return m
}

func (m *Matrix) index(row int, col int) int {
	return row*m.nrCols + col
}

func (m *Matrix) within(row int, col int) bool {
	var inRows bool = 0 <= row && row < m.nrRows
	var inCols bool = 0 <= col && col < m.nrCols
	return inRows && inCols
}

func (m *Matrix) Set(row int, col int, v int) error {
	if m.within(row, col) {
		index := m.index(row, col)
		m.values[index] = v
		return nil
	}
	return errors.New("Set of coordinate outside matrix")
}

func (m *Matrix) Get(row int, col int) (int, error) {
	if m.within(row, col) {
		index := m.index(row, col)
		return m.values[index], nil
	}
	return 0, errors.New("Get of coordinate outside matrix")
}

func (m *Matrix) AppendRow(row []int) error {
	if len(row) == m.nrCols {
		m.values = append(m.values, row...)
		m.nrRows++
		return nil
	}
	return errors.New("appenden row length does not match column length")
}

func (m Matrix) String() string {
	var representation string = ""
	for index, element := range m.values {
		representation = representation + strconv.Itoa(element)
		if (index+1)%m.nrCols == 0 {
			representation = representation + "\n"
		}
	}
	return representation
}

func (m Matrix) Inbounds(position Vector2d) bool {
	return 0 <= position.x && position.x < m.nrRows &&
		0 <= position.y && position.y < m.nrCols
}

func (m Matrix) GetNrRows() int {
	return m.nrRows
}

func (m Matrix) GetNrCols() int {
	return m.nrCols
}

func (m Matrix) SumElements() int {
	sum := 0
	for _, el := range m.values {
		sum += el
	}
	return sum
}
