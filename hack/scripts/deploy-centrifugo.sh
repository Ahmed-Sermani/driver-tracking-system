#!/bin/bash
set -e

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "Deploying Centrifugal using Helm..."
helm repo add centrifugal https://centrifugal.github.io/helm-charts
helm repo update
helm upgrade centrifugo centrifugal/centrifugo -f "$parent_path/centrifugo-values.yaml" -n messaging --create-namespace --install --wait
