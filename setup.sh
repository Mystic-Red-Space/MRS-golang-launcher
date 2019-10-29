#!/usr/bin/env bash

GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

go get -u -v github.com/asticode/go-astilectron
go get -u -v github.com/asticode/go-astilectron-bootstrap
go get -u -v github.com/asticode/go-astilectron-bundler/...
go get -u -v github.com/asticode/go-astilog
go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
