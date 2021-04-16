# Prometheus and Grafana Project

In this project there is an API that will expose its metrics using Prometheus client SDK, and this will be used by Prometheus container to collect metrics.
Also there is a Grafana container that will use Prometheus as datasource and create dashboards using these metrics, everything automatically.

## Up and Running

```bash
docker-compose up -d
```

## Testing

```bash
# calling prometheus API directly
curl --silent http://localhost:9090/metrics

# calling metrics of the API which is using prometheus client
curl --silent http://localhost:9000/metrics

# calling metrics of the API searching for the custom metric
curl --silent http://localhost:9000/metrics \
  | grep custom_request_healthcheck_counter
# custom_request_healthcheck_counter 0

# incrementing the counter of this custom metric by calling the healthcheck endpoint
curl --silent http://localhost:9000/healthcheck

# and then calling metrics of the API again
curl --silent http://localhost:9000/metrics \
  | grep custom_request_healthcheck_counter
# custom_request_healthcheck_counter 1
```
