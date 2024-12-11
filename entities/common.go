package entities

const ErrorCodeInternalServerError = "500"

type ErrorResponse struct {
	Code  string `json:"code"`  // Response Code
	Error string `json:"error"` // Response Error
}
