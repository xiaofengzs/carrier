#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./carrier/vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

../vendor/k8s.io/code-generator/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  carrier/pkg/generated\
  carrier/pkg/apis \
  xiaofeng:v1 \
  --output-base $(pwd)/../.. \
  --go-header-file $(pwd)/boilerplate.go.txt