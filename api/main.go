package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsHealthcheckCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Subsystem: "custom",
		Name: "custom_request_healthcheck_counter",
		Help: "The total number of requests made to healthcheck endpoint",
	})
)

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

		requestsHealthcheckCounter.Inc()

		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(statusCode)
		writter.Write(responseJSONBytes)
	}
}

func main() {
	addr := "0.0.0.0:9000"

	// register counter
	prometheus.Register(requestsHealthcheckCounter)

	// endpoint for healthcheck
	http.HandleFunc("/healthcheck", returnHTTPResponse(http.StatusOK, "OK"))

	// endpoint for prometheus handler
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Listening at %s", addr)

	http.ListenAndServe(addr, nil)
}
