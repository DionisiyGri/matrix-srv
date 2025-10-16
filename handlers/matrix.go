package handlers

import (
	"fmt"
	"league-matrix/internal/matrix"
	"league-matrix/parser"
	"net/http"
	"strings"
)

type matrixHandler struct {
	parser    parser.Parser
	matrixSvc matrix.Matrixer
}

func New(parser parser.Parser,
	matrixSvc matrix.Matrixer,
) *matrixHandler {
	return &matrixHandler{
		parser:    parser,
		matrixSvc: matrixSvc,
	}
}

// Return the matrix as a string in matrix format
func (mh *matrixHandler) Echo(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := mh.matrixSvc.Echo(body)
	fmt.Fprint(w, res)
}

// Return the matrix as a string in matrix format where the columns and rows are inverted
func (mh *matrixHandler) Invert(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inverted := mh.matrixSvc.Invert(body)

	var res string
	for _, row := range inverted {
		res = fmt.Sprintf("%s%s\n", res, strings.Join(row, ","))
	}
	fmt.Fprint(w, res)
}

// Return the matrix as a 1 line string, with values separated by commas
func (mh *matrixHandler) Flatten(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	flattened := mh.matrixSvc.Flatten(body)
	fmt.Fprintf(w, strings.Join(flattened, ","))
}

// Return the sum of the integers in the matrix
func (mh *matrixHandler) Sum(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum, err := mh.matrixSvc.Sum(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d", sum)
}

// Return the product of the integers in the matrix
func (mh *matrixHandler) Multiply(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	multiply, err := mh.matrixSvc.Multiply(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d", multiply)
}
