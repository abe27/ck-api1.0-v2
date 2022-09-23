package models

type Response struct {
	Success bool        `json:"success" default:"false"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
