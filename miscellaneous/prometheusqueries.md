# Prometheus sample queries

```bash
# promhttp
promhttp_metric_handler_requests_total{job="api"}
promhttp_metric_handler_requests_total{job="api",code="200"}

rate(promhttp_metric_handler_requests_total{job="api"}[1m])

sum(rate(promhttp_metric_handler_requests_total{job="api"}[1m]))


# http_request_healthcheck_count
custom_http_request_healthcheck_count{status="success"}

# taxa de erros nas requests
sum(rate(custom_http_request_healthcheck_count{status="error"}[1m])) / sum(rate(custom_http_request_healthcheck_count{}[1m]))


# http_request_healthcheck_duration_seconds
custom_http_request_healthcheck_duration_seconds_sum{status="success"}

custom_http_request_healthcheck_duration_seconds_count{status="success"}

custom_http_request_healthcheck_duration_seconds_bucket{status="success"}

# average duration
rate(custom_http_request_healthcheck_duration_seconds_sum{status="success"}[1m]) / rate(custom_http_request_healthcheck_duration_seconds_count{status="success"}[1m])

histogram_quantile(0.5, rate(custom_http_request_healthcheck_duration_seconds_bucket{status="success"}[1m]))
```
