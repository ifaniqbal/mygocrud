package model

type Request struct {
	Page     int    `json:"page,omitempty" form:"page"`
	PageSize int    `json:"page_size,omitempty" form:"page_size"`
	Filter   string `json:"filter,omitempty" form:"filter"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
