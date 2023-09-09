package matrixp

type Vector2d struct {
	x int
	y int
}

func add(v1 Vector2d, v2 Vector2d) Vector2d {
	return Vector2d{v1.x + v2.x, v1.y + v2.y}
}

type Directional struct {
	position  Vector2d
	direction Vector2d
	matrix    *Matrix
}

func (i *Directional) HasNext() bool {
	return i.matrix.Inbounds(i.position)
}

func (i *Directional) GetNext() (int, int, int) {
	v := i.position
	i.position = add(i.position, i.direction)
	value, _ := i.matrix.Get(v.x, v.y)
	return v.x, v.y, value
}

func InitLineForward(row int, matrix *Matrix) Directional {
	return InitRight(row, 0, matrix)
}

func InitLineBackward(row int, matrix *Matrix) Directional {
	return InitLeft(row, matrix.nrCols-1, matrix)
}

func InitColumnForeward(col int, matrix *Matrix) Directional {
	return InitDown(0, col, matrix)
}

func InitColumnBackward(col int, matrix *Matrix) Directional {
	return InitUp(matrix.nrRows-1, col, matrix)
}

func InitUp(row int, col int, matrix *Matrix) Directional {
	start := Vector2d{row, col}
	direction := Vector2d{-1, 0}
	return Directional{position: start, direction: direction, matrix: matrix}
}

func InitRight(row int, col int, m *Matrix) Directional {
	start := Vector2d{row, col}
	direction := Vector2d{0, 1}
	return Directional{start, direction, m}
}

func InitDown(row int, col int, m *Matrix) Directional {
	start := Vector2d{row, col}
	l := Vector2d{1, 0}
	return Directional{start, l, m}
}

func InitLeft(row int, col int, m *Matrix) Directional {
	start := Vector2d{row, col}
	l := Vector2d{0, -1}
	return Directional{start, l, m}
}
