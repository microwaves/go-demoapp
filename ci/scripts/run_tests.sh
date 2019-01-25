#!/bin/sh

set -e -u -x

export GOPATH=$PWD/git-go-demoapp
go test
