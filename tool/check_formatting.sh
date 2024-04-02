#!/usr/bin/env bash

set -e

if [ "$(go fmt './...' | wc -l)" -gt 0 ]; then
    echo "Some files are not formatted. Please run 'go fmt ./...' and commit the changes."
    exit 1
fi
