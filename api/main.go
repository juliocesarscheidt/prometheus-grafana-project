package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/http"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	healthcheckEndpointCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "custom",
			Name:      "http_request_healthcheck_count",
			Help:      "The total number of requests made to healthcheck endpoint",
		},
		[]string{"status"},
	)

	healthcheckEndpointLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "custom",
			Name: "http_request_healthcheck_duration_seconds",
			Help: "Latency of healthcheck endpoint requests in seconds",
		},
		[]string{"status"},
	)
)

func init() {
	// must register counter on init
	prometheus.MustRegister(healthcheckEndpointCounter)
	prometheus.MustRegister(healthcheckEndpointLatency)
}

func buildJSONResponse(statusCode int, message string) ([]byte, error) {
	var responseHTTP = make(map[string]interface{})

	responseHTTP["statusCode"] = statusCode
	responseHTTP["data"] = message

	response, err := json.Marshal(responseHTTP)
	if err != nil {
		return nil, err
	}

	return []byte(string(response)), nil
}

func returnHTTPResponse(statusCode int, message string) http.HandlerFunc {
	return func(writter http.ResponseWriter, req *http.Request) {
		responseJSONBytes, _ := buildJSONResponse(statusCode, message)

		var status string
		// counter for prometheus
		defer func() {
			healthcheckEndpointCounter.WithLabelValues(status).Inc()
		}()
		// latency timer for prometheus
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(_time float64) {
			healthcheckEndpointLatency.WithLabelValues(status).Observe(_time)
		}))
		defer func() {
			timer.ObserveDuration()
		}()

		// sleep from 0 to 2 secs randomly
		duration := (rand.Intn(2 - 0) + 0)
		time.Sleep(time.Duration(duration) * time.Second)

		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(statusCode)

		status = "success"

		writter.Write(responseJSONBytes)
	}
}

func main() {
	addr := "0.0.0.0:9000"

	// endpoint for healthcheck
	http.HandleFunc("/healthcheck", returnHTTPResponse(http.StatusOK, "OK"))

	// endpoint for prometheus handler
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Listening at %s", addr)

	http.ListenAndServe(addr, nil)
}
