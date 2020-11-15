#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
cd "$BASEDIR/../.." || exit

BUILD_TARGET_DIR="./dist"

rm -rf $BUILD_TARGET_DIR || true
mkdir -p ${BUILD_TARGET_DIR}

go build \
    -o "${BUILD_TARGET_DIR}/inmemory" \
    -gcflags="all=-N -l" \
    ./main.go

cp ./contrib/config.yml ${BUILD_TARGET_DIR}/config.yml || true
cp ./contrib/swagger/openapi.json ${BUILD_TARGET_DIR}/openapi.json || true




