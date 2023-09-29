#!/bin/bash
set -e

case $(uname -s) in
  Darwin|Linux)
    ;;
  *)
    echo "This script only works on Mac or Linux"
    exit 1
    ;;
esac

check_command() {
  local command=$1
  if ! command -v $command &> /dev/null; then
    echo "$command is not installed"
    exit 1
  fi
}

check_command docker
check_command kubectl
check_command kind
check_command helm

# Check if docker is running
if ! docker info &> /dev/null; then
  echo "Docker is not running. Please start it"
  exit 1
fi

cluster_config=$(mktemp)
cat > $cluster_config <<'EOF'
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
    - |
      kind: InitConfiguration
      nodeRegistration:
        kubeletExtraArgs:
          node-labels: "ingress-ready=true"        
  extraPortMappings:
    - containerPort: 80
      hostPort: 80
      protocol: TCP
- role: worker
EOF

echo "Creating a K8S cluster with two nodes."
kind create cluster --name driver-tracking-system --image kindest/node:v1.24.6 --config=$cluster_config

rm $cluster_config

echo "The cluster is created successfully"
kubectl cluster-info --context kind-driver-tracking-system
kubectl get nodes -o wide --context kind-driver-tracking-system 

