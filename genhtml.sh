#!/usr/bin/env bash

set -e

dart ./packages/astout/bin/astout.dart $1 > asthtml.html
open asthtml.html
