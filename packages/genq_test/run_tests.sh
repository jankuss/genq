#!/usr/bin/env bash
set -e

if [ -z "$GENQ_PATH" ]; then
  GENQ_PATH=genq
fi

$GENQ_PATH

dart test
