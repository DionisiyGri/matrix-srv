package parser

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Parser interface {
	ParseCSV(file io.Reader) ([][]string, error)
}

type parser struct{}

func New() Parser {
	return &parser{}
}

func (p *parser) ParseCSV(file io.Reader) ([][]string, error) {
	c, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("invalid file content: %w", err)
	}
	return c, nil
}
