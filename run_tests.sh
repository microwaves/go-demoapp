#!/bin/bash

set -e -u -x

export GOPATH=$PWD
go test
