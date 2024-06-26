package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"go.guide/mul-grpc-service/pb"
	"io/ioutil"
	"net/http"
	"path"
)

func _Decode_Grpc_Mul_Request(ctx context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func _Encode_Grpc_Mul_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	return resp, nil
}

func _Decode_Http_Mul_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.MulRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Http_Mul_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp pb.MulResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_Http_Mul_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "add")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Http_Mul_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
