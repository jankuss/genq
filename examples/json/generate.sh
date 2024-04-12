#!/usr/bin/env bash

if [ -z "$GENQ_PATH" ]; then
  GENQ_PATH=genq
fi

# Run command under GENQ_PATH
$GENQ_PATH
