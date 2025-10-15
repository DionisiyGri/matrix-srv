package handlers

import (
	"fmt"
	"net/http"
)

func (mh *matrixHandler) parseRequest(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %w", err)
	}
	defer file.Close()

	records, err := mh.parser.ParseCSV(file)
	if err != nil {
		return nil, fmt.Errorf("invalid data format: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("error: file is empty")
	}

	return records, nil
}
