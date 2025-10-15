package matrix

import (
	"fmt"
	"strings"
)

type Matrixer interface {
	Echo(input [][]string) string
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
