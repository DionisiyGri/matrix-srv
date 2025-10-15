package matrix

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Matrixer interface {
	Echo(input [][]string) string
	Invert(input [][]string) [][]string
	Flatten(input [][]string) []string
	Sum(input [][]string) (int, error)
	Multiply(input [][]string) (int, error)
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

func (m csvMatrix) Flatten(input [][]string) []string {
	var res []string
	for _, row := range input {
		res = append(res, row...)
	}
	return res
}

func (m csvMatrix) Sum(input [][]string) (int, error) {
	var resSum int
	for _, row := range input {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("cannot convert string to int - %v", err)
				return 0, errors.New("invalid value")
			}
			resSum += num
		}
	}
	return resSum, nil
}

func (m csvMatrix) Multiply(input [][]string) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}

	resMult := 1
	for _, row := range input {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("cannot convert string to int - %v", err)
				return 0, errors.New("invalid value")
			}
			resMult *= num
		}
	}
	return resMult, nil
}
