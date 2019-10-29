#!/usr/bin/env bash

export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

bin/astilectron-bundler -v
