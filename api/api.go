package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Only success code like 200 will be returned
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error code like 404, 503, etc.
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, error.Error(), http.StatusBadRequest)
	}
	InternalErrorHanlder = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}
)
