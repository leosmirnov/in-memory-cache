#!/usr/bin/env bash

BASEDIR=$(dirname "$0")

PROJECT_ROOT=$(
    cd "$BASEDIR/../.." || exit
    pwd -P
)

docker build \
      -f "${PROJECT_ROOT}"/contrib/docker/Dockerfile \
      -t leosmirnov/inmemory \
      .

docker run \
    --rm \
    -v "${PROJECT_ROOT}"/contrib/config.yml:/dist/config.yml \
    -p 8080:8080 \
    --name inmemory-storage \
    leosmirnov/inmemory