package parser

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Parser interface {
	Parse(file io.Reader) ([][]string, error)
}

type fileParser struct{}

func New() Parser {
	return &fileParser{}
}

func (fp *fileParser) Parse(file io.Reader) ([][]string, error) {
	c, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("invalid file content: %w", err)
	}
	return c, nil
}
