#!/bin/bash
set -e

echo "Deploying nginx-ingress controller..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

# Wait for ingress controller to be ready
while [[ $(kubectl get pods -n ingress-nginx --selector=app.kubernetes.io/component=controller -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do
  echo "Waiting for ingress controller to be ready..."
  sleep 5
done
