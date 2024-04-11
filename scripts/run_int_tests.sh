#!/usr/bin/env bash
set -e

# Is GENQ_PATH set?
if [ -z "$GENQ_PATH" ]; then
  rm tool/genq || true
  (cd tool && go build .)
  GENQ_PATH=../genq
fi

(cd tool/int_test && GENQ_PATH=$GENQ_PATH go test -p 1)
