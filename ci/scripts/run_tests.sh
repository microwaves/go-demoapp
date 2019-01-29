#!/bin/sh

set -e -u -x
cd git && go test -a -tags netgo -ldflags '-w'
