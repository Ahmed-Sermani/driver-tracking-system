#!/bin/bash
set -e

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "Deploying Centrifugal using Helm..."
helm repo add stakater https://stakater.github.io/stakater-charts
helm repo update
helm upgrade monolith stakater/application -f "$parent_path/app-values.yaml" --namespace app --create-namespace --install --wait
