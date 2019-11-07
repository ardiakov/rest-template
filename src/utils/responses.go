package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

func SuccessResponse(data []string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status: "success",
		Data:   data,
	}

	json, _ := json.Marshal(response)

	w.Write(json)
}

func ErrorResponse(data []string, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := Response{
		Status: "error",
		Data:   data,
	}

	json, _ := json.Marshal(response)
	w.Write(json)
}
