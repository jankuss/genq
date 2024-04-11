#!/usr/bin/env bash

rm tool/int_test/genq || true

set -e
(cd tool && go build .)

mv tool/genq tool/int_test

PATH=./tool/int_test:$PATH

(cd tool/int_test && go test -p 1 -run TestStdin)
