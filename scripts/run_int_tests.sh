#!/usr/bin/env bash
set -e

rm tool/genq || true
(cd tool && go build .)
(cd tool/int_test && PATH=./tool:$PATH go test -p 1 -run TestStdin)
