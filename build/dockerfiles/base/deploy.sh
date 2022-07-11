#!/usr/bin/env bash
set -eo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"

image=registry.erda.cloud/erda/erda-base:$(date +"%Y%m%d")

docker buildx build --platform linux/amd64,linux/arm64 -t ${image} --push . -f Dockerfile

echo "action meta: erda-base=$image"
