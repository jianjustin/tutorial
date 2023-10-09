package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"jianjustin/sub-grpc-service/transport"
	"net/http"
	"path"
)

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

func _Decode_Sub_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.SubRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Sub_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.SubResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_Sub_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "sub")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_SubAfterAdd_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "addAfterMul")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Sub_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
