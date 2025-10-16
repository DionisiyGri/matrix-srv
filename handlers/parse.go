package handlers

import (
	"fmt"
	"log"
	"net/http"
)

const maxUploadSize = 5 << 20 // 5 MB

// parseRequest reads a CSV file from a multipart form request and returns its records
func (mh *matrixHandler) parseRequest(w http.ResponseWriter, r *http.Request) ([][]string, error) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		log.Printf("invalid file size: %v", err)
		return nil, fmt.Errorf("invalid file size")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		return nil, fmt.Errorf("failed to fetch file")
	}
	defer file.Close()

	records, err := mh.parser.ParseCSV(file)
	if err != nil {
		log.Printf("error parsing file: %v", err)
		return nil, fmt.Errorf("invalid data format")
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return records, nil
}
