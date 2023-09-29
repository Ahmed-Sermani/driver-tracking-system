#!/bin/bash
set -e

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "Deploying Redis using Helm..."
helm install redis oci://registry-1.docker.io/bitnamicharts/redis -f "$parent_path/redis-values.yaml" -n redis --create-namespace 
