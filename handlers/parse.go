package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// parseRequest reads a CSV file from a multipart form request and returns its records
func (mh *matrixHandler) parseRequest(r *http.Request) ([][]string, error) {
	err := r.ParseMultipartForm(1 << 0) //5mb max
	if err != nil {
		log.Printf("fnvalid file size: %w", err)
		return nil, fmt.Errorf("invalid file size")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %w", err)
		return nil, fmt.Errorf("failed to fetch file")
	}
	defer file.Close()

	records, err := mh.parser.ParseCSV(file)
	if err != nil {
		log.Printf("error parsing file: %w", err)
		return nil, fmt.Errorf("invalid data format")
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return records, nil
}
