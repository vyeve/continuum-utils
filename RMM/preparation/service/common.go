package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	partnerID  = "partnerID"
	endpointID = "endpointID"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json; charset=utf-8"
)

// GetPartnerID extracts PartnerID from HTTP Request
func GetPartnerID(r *http.Request) string {
	return mux.Vars(r)[partnerID]
}

// GetEndpointID extracts GetEndpointID from HTTP Request
func GetEndpointID(r *http.Request) string {
	return mux.Vars(r)[endpointID]
}

//Error contains the message about error
type Error struct {
	Message string `json:"message"`
}

//ErrorMessage contains the error
type ErrorMessage struct {
	Error Error `json:"error"`
}

// renderJSON is used for rendering JSON response
func renderJSON(w http.ResponseWriter, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		SendInternalServerError(w, fmt.Sprintf("helpers.renderJSON: %s", err))
		return
	}
	render(w, status, data)
}

// RenderJSON is used for rendering JSON response body with appropriate headers
func RenderJSON(w http.ResponseWriter, response interface{}) {
	renderJSON(w, http.StatusOK, response)
}

func render(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(code)
	w.Write(response)
}

var createError = func(msg string) interface{} {
	return ErrorMessage{Error{msg}}
}

// SendError writes a defined string as an error message
// with appropriate headers to the HTTP response
func SendError(w http.ResponseWriter, code int, message string) {
	if message == "" {
		message = http.StatusText(code)
	}
	data, _ := json.Marshal(createError(message))
	render(w, code, data)
}

// SendInternalServerError sends Internal Server Error Status and logs an error if it exists
func SendInternalServerError(w http.ResponseWriter, message string) {
	SendError(w, http.StatusInternalServerError, message)
}
