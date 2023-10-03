```bash
docker build -f Dockerfile.prod --tag arfurlaneto/go-expert-docker .
docker push arfurlaneto/go-expert-docker
```

```bash
kind create cluster --name=goexpert
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl port-forward service/serversvc 8080:8080
```
