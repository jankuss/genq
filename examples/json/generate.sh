#!/usr/bin/env bash

(cd ../../tool && go build -o ../examples/json)
./genq
