package models

// Result is an error structure sent to client if he wrong
type Result struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    string `json:"code"`
}
