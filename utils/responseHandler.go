package utils

import (
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Status   int                    `json:"status"`
	Message  string                 `json:"message"`
	Response map[string]interface{} `json:"response"`
}

func ResponseEncoders(rs http.ResponseWriter, responseStatus int, responseData any, message string) {
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(responseStatus)
	response := UserResponse{Status: responseStatus, Message: message, Response: map[string]interface{}{"data": responseData}}
	json.NewEncoder(rs).Encode(response)
}
