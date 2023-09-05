package matrix

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	var m Matrix = Init(2, 4)
	standardV, error1 := m.Get(0, 0)
	if error1 != nil {
		t.Errorf("error at get (0,0)")
	}
	if standardV != 0 {
		t.Errorf("not standard value at (0,0)")
	}
	error2 := m.Set(1, 2, 10)
	if error2 != nil {
		t.Errorf("error with set within matrix")
	}
	v, error3 := m.Get(1, 2)
	if error3 != nil {
		t.Errorf("error when ger within matrix")
	}
	if v != 10 {
		t.Errorf("set get loop failed")
	}
}

func TestGetError(t *testing.T) {
	var m Matrix = Init(0, 0) //empty matrix
	v, error := m.Get(0, 0)
	if v != 0 || error == nil {
		t.Errorf("no error raised when accesing an empty matrix")
	}
}

func TestAppend(t *testing.T) {
	var m Matrix = Init(1, 2)
	m.AppendRow([]int{1, 2})

}
