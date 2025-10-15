package handlers

import (
	"fmt"
	"league-matrix/internal/matrix"
	"league-matrix/parser"
	"net/http"
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

func (mh *matrixHandler) Echo(w http.ResponseWriter, r *http.Request) {
	body, err := mh.parseRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := mh.matrixSvc.Echo(body)
	fmt.Fprint(w, res)
}
