package response

import (
	"encoding/json"
	"finance-solver.com/models"
	"log"
	"net/http"
)

const (
	StatusOK     = "ok"
	GenericError = "error"
)

type JsonResponse struct {
	StatusCode int            `json:"-"`
	Status     string         `json:"status"`
	Message    string         `json:"message,omitempty"`
	Data       *[]interface{} `json:"data,omitempty"`
}

func (j *JsonResponse) Success(w http.ResponseWriter, message string, statusCode int) {
	j.StatusCode = statusCode
	j.Status = StatusOK
	j.Message = message
	j.Render(w)
}

func (j *JsonResponse) OkData(w http.ResponseWriter, data *[]interface{}, statusCode int) {
	j.StatusCode = statusCode
	j.Status = StatusOK
	j.Data = data
	j.Render(w)
}

func (j *JsonResponse) Error(w http.ResponseWriter, err error, statusCode int) {
	j.StatusCode = statusCode
	j.Status = GenericError

	if publicError, ok := err.(models.PublicError); ok {
		j.Message = publicError.Public()
	} else {
		j.Message = err.Error()
	}

	j.Render(w)
}

func (j *JsonResponse) ErrorWithMessage(w http.ResponseWriter, msg string, statusCode int) {
	j.StatusCode = statusCode
	j.Status = GenericError
	j.Message = msg
	j.Render(w)
}

func (j *JsonResponse) Render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(j.StatusCode)
	enc := json.NewEncoder(w)
	err := enc.Encode(j)
	if err != nil {
		// todo: handle this in the response also
		log.Println(err)
	}
}

//
//
//func NewResponse() *JsonResponse {
//	return &JsonResponse{
//		StatusCode: 0,
//		Status:     "",
//		Message:    "",
//		Data:       nil,
//	}
//}
//
//func NewError() *JsonResponse {
//	return &JsonResponse{
//		StatusCode: 0,
//		Status:     "",
//		Message:    "",
//		Data:       nil,
//	}
//}

//
//func rOk() {
//
//}
//
//func rError() {
//
//}
//
//func renderJSON(w http.ResponseWriter, data []interface{}, statusCode int, message string) {
//
//	j := JsonResponse{
//		StatusCode: statusCode,
//		Status:     "ok",
//		Message:    message,
//		Data:       data,
//	}
//	j.Render(w)
//}

//func renderError(w http.ResponseWriter, err error, status int) {
//	renderErrorWithMessage(w, err, status, "error")
//}

//func renderSuccess(w http.ResponseWriter) {
//	renderJSON(w, struct {
//		Message string `json:"message"`
//	}{Message: "success"}, http.StatusOK)
//}
//
//func renderErrorWithMessage(w http.ResponseWriter, err error, status int, msg string) {
//	e := jsonError{
//		Message: msg,
//		Error:   err.Error(),
//	}
//	renderJSON(w, e, status)
//}

//func renderError(w http.ResponseWriter, err error, status int) {
//	renderErrorWithMessage(w, err, status, "error")
//}
