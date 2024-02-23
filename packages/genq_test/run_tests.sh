#!/usr/bin/env bash
set -e

(cd ../../tool && go build .)
cp ../../tool/genq .

./genq

dart test
