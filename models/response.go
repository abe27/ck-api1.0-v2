package models

type Response struct {
	Success     bool        `json:"success,omitempty" default:"false"`
	TestMessage string      `json:"testMessage,omitempty"`
	TestObject  interface{} `json:"test_object,omitempty"`
	Message     string      `json:"message,omitempty"`
	Error       interface{} `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Page        int         `json:"page,omitempty"`
	TotalPages  int         `json:"totalPages,omitempty"`
}
