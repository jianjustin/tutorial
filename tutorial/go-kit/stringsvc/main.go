package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"go.guide/tutorial/go-kit/stringsvc/logging"
	"go.guide/tutorial/go-kit/stringsvc/service"
	"go.guide/tutorial/go-kit/stringsvc/transport"
	"net/http"
	"os"
)

func main() {
	//svc := service.StringServiceImpl{}
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = service.StringServiceImpl{}
	svc = logging.LoggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		transport.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transport.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
