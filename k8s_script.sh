#!/bin/bash
set -e

echo "⏳ Enabling ingress addon..."
minikube addons enable ingress

echo "⏳ Applying Deployment, Service, and Ingress..."
kubectl apply -f k8s/hello-app.yaml

echo "⏳ Waiting for pods to be ready..."
kubectl wait --for=condition=ready pod -l app=hello-app --timeout=120s

# Get minikube IP
MINIKUBE_IP=$(minikube ip)
echo "🏠 Minikube IP: $MINIKUBE_IP"

# Map /etc/hosts
echo "⏳ Adding /etc/hosts entry..."
if ! grep -q "hello-app.local" /etc/hosts; then
  echo "$MINIKUBE_IP hello-app.local" | sudo tee -a /etc/hosts
fi

echo "⏳ Verifying service connectivity using curl pod..."
kubectl run curlpod \
  --rm -it \
  --image=curlimages/curl:7.85.0 \
  --restart=Never \
  -- curl -s http://hello-app-service.default.svc.cluster.local:80/

echo "⏳ Verifying ingress connectivity..."
curl -v http://hello-app.local/
