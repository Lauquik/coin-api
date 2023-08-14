package api

import (
	"encoding/json"
	"net/http"
)

type Checkbalacne struct {
	Username string
}
type CheckbalacneResp struct {
	Balance    int64
	StatusCode int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("contetn-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestError = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHander = func(w http.ResponseWriter) {
		writeError(w, "some internal error occured", http.StatusInternalServerError)
	}
)
