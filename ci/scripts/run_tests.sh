#!/bin/sh

set -e -u -x

export GOPATH=$PWD
go test
