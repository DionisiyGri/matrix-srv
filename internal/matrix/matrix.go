package matrix

import (
	"fmt"
	"strings"
)

type Matrixer interface {
	Echo(input [][]string) string
	Invert(input [][]string) [][]string
}

type csvMatrix struct{}

func NewCSVMatrixer() Matrixer {
	return csvMatrix{}
}

func (m csvMatrix) Echo(input [][]string) string {
	var res string
	for _, row := range input {
		res = fmt.Sprintf("%s%s\n", res, strings.Join(row, ","))
	}
	return res
}

func (m csvMatrix) Invert(input [][]string) [][]string {
	numRows := len(input)
	numCols := len(input[0])

	res := make([][]string, numCols)

	for i := 0; i < numCols; i++ {
		res[i] = make([]string, numRows)
		for j := 0; j < numRows; j++ {
			res[i][j] = input[j][i]
		}
	}
	return res
}
