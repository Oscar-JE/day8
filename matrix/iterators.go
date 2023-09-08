package matrix

type MatrixIterator interface {
	hasNext() bool
	getNext() int
}

type LineForward struct {
	index  int
	row    int
	matrix *Matrix
}

func InitLineFoeward(row int, matrix *Matrix) LineForward {
	index := 0
	return LineForward{index: index, row: row, matrix: matrix}
}

func (i *LineForward) hasNext() bool {
	return i.index < i.matrix.nrCols
}

func (i *LineForward) getNext() int {
	if i.hasNext() {
		value, _ := i.matrix.Get(i.row, i.index)
		i.index++
		return value
	}
	return -1
}
