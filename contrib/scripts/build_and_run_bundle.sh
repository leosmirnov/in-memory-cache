#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
cd "$BASEDIR" || exit

bash ./build_bundle.sh "$@" || exit 1
bash ./run_bundle.sh "$@" || exit 1
