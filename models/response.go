package models

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
