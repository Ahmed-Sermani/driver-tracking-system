#!/bin/bash
set -e

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
echo "Deploying Postgres..."

helm install postgres oci://registry-1.docker.io/bitnamicharts/postgresql -f "$parent_path/postgres-values.yaml" -n db --create-namespace
