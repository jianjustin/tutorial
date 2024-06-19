package model

import (
	"bytes"
	"context"
	"encoding/json"
	"go.guide/add-service/model"
	"io/ioutil"
	"net/http"
)

type MulRequest struct {
	A int `json:"a"`
}

type MulResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

func DecodeMulRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request MulRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func DecodeAddResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response model.AddResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
