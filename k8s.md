# Running on K8s

## Create monitoring namespace

```bash
kubectl create ns monitoring
```

## Install prometheus chart

> https://artifacthub.io/packages/helm/bitnami/prometheus-operator
> https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm upgrade -i prometheus-operator bitnami/prometheus-operator -n monitoring

helm ls -n monitoring

# port forward for prometheus UI access
kubectl port-forward -n monitoring svc/prometheus-operator-prometheus 9090:9090 --address 0.0.0.0 &
```

## Install grafana chart

> https://bitnami.com/stack/grafana-operator/helm
> https://hub.kubeapps.com/charts/bitnami/grafana-operator
> https://docs.bitnami.com/tutorials/manage-multiple-grafana-operator

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm upgrade -i grafana-operator bitnami/grafana-operator -n monitoring --set grafana.config.security.admin_user="admin" --set grafana.config.security.admin_password="admin"

helm ls -n monitoring

# port forward for prometheus UI access
kubectl port-forward -n monitoring service/grafana-service 3000:3000 --address 0.0.0.0 &
```

# Running API

```bash
docker image build --tag juliocesarmidia/api:v1.0.0 "${PWD}/api/"
# docker container run -d --rm --name api --publish 9000:9000 juliocesarmidia/api:v1.0.0
# docker container logs -f api

# curl -X GET http://localhost:9000/metrics

# docker container rm -f api


kubectl apply -f api.yaml
kubectl get pod,deploy,svc,servicemonitor -n default


SERVICE_IP="$(kubectl get svc -l component=api --no-headers | tr -s ' ' ' ' | cut -d' ' -f3)"
echo "${SERVICE_IP}"

curl -X GET -L "http://${SERVICE_IP}/healthcheck"
curl -X GET -L "http://${SERVICE_IP}/metrics"

# kubectl delete -f api.yaml
```
