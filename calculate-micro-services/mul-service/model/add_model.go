package model

type AddRequest struct {
	A int `json:"a"`
}

type AddResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}
