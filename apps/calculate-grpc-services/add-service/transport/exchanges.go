package transport

type AddRequest struct {
	A int64 `json:"a"`
}

type AddResponse struct {
	V int64 `json:"v"`
}
