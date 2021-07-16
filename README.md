# Prometheus and Grafana Project

In this project there is an API that will expose its metrics using Prometheus client SDK, and this will be used by Prometheus container to collect metrics.
Also there is a Grafana container that will use Prometheus as datasource and create dashboards using these metrics, everything automatically.

## Up and Running

```bash
docker-compose up -d
```

Then access Grafana on: <http://localhost:3000>

> Grafana default user: admin/admin
> Prometheus default user: admin/admin

> Generate encrypted password

```bash
htpasswd -nb -B admin admin
admin:$2y$05$aShml8bxjrquQLKe0H6VE.M2lG4nVjasrYheS2LOukWOV0zTlPcSS
# used options:
-n  Don't update file; display results on stdout.
-m  Force MD5 encryption of the password (default).
-B  Force bcrypt encryption of the password (very secure)
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
declare -i COUNTER
let "COUNTER+=0"
while true; do
  test $COUNTER -ge 50 && break
  curl --silent "http://localhost:9000/healthcheck"
  let "COUNTER+=1"
done

# and then calling metrics of the API again
curl --silent http://localhost:9000/metrics \
  | grep custom_request_healthcheck_counter
# custom_request_healthcheck_counter 1
```
