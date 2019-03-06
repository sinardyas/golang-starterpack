package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) ResponseHandling(res http.ResponseWriter, status int, success bool, message string, data interface{}) {
	result := Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	}

	parsedResult, _ := json.Marshal(result)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(parsedResult)
}
