package model

type MulRequest struct {
	A int `json:"a"`
}

type MulResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}
