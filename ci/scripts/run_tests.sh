#!/bin/sh

set -e -u -x
cd git-go-demoapp && go test -a -tags netgo -ldflags '-w'
