#!/usr/bin/env bash
set -e

rm tool/genq || true
(cd tool && go build .)
PATH=./tool:$PATH
whereis genq
(cd tool/int_test && go test -p 1 -run TestStdin)
