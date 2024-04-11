#!/usr/bin/env bash
set -e

rm tool/genq || true
(cd tool && go build .)
(cd tool/int_test && GENQ_PATH=../genq go test -p 1 -run TestStdin)
