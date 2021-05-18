package matrix

import (
	"strings"
	"errors"
	"strconv"
)

type Matrix struct {
	data [][]int
}

func New(input string) (*Matrix, error) {
	lines := strings.Split(input, "\n")
	var lenColsPrev int = -1

	var data [][]int

	for _, l := range lines {
		l = strings.TrimSpace(l)
		items := strings.Split(l, " ")
		lenColsNew := len(items)

		if lenColsPrev != -1 && lenColsNew != lenColsPrev {
			return nil, errors.New("malformat")
		}

		lenColsPrev = lenColsNew

		var row []int
		for _, item := range items {
			i, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		data = append(data, row)
	}

	return &Matrix{data}, nil
}

func (m *Matrix)Rows() [][]int {
	var res [][]int
	for _, rows := range m.data {
		var row []int
		for _, item := range rows {
			row = append(row, item)
		}
		res = append(res, row)
	}
	return res
}

func (m *Matrix)Cols() [][]int {
	xl := len(m.data[0])
    yl := len(m.data)

	var res [][]int

	for i := 0; i < xl; i++ {
		var col []int

		for j := 0; j < yl; j++ {
			col = append(col, m.data[j][i])
		}
		res = append(res, col)
	}

	return res
}

func (m *Matrix)Set(r int, c int, v int) bool {

	xl := len(m.data[0])
    yl := len(m.data)

	if r < 0 || r >= yl || c < 0 || c >= xl {
		return false
	}

	m.data[r][c] = v
	return true
}
