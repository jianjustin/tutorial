package transporthttp

import (
	"fmt"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jianjustin/sub-grpc-service/transport"
	"math/rand"
	http1 "net/http"
	"time"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/sub").Handler(
		http.NewServer(
			endpoints.SubEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "sub", logger)))...))
	mux.Methods("POST").Path("/subAfterAdd").Handler(
		http.NewServer(
			endpoints.SubAfterAddEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "subAfterAdd", logger)))...))

	requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of the HTTP request durations in seconds.",
		Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	})

	// Create non-global registry.
	registry := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		requestDurations,
	)

	go func() {
		for {
			// Record fictional latency.
			now := time.Now()
			requestDurations.(prometheus.ExemplarObserver).ObserveWithExemplar(
				time.Since(now).Seconds(), prometheus.Labels{"dummyID": fmt.Sprint(rand.Intn(100000))},
			)
			time.Sleep(600 * time.Millisecond)
		}
	}()

	// Expose /metrics HTTP endpoint using the created custom registry.
	mux.Methods("GET").Path("/metrics").Handler(promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		}))
	return mux
}
