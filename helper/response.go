package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"message"`
}

func (r *Response) ResponseHandling(res http.ResponseWriter, status int, success bool, data interface{}) {
	result := Response{
		Status:  status,
		Success: success,
		Data:    data,
	}

	parsedResult, _ := json.Marshal(result)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(parsedResult)
}
