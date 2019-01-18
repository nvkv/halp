#!/usr/bin/env bash

export GO111MODULE=on

#go get
go test github.com/hellohippo/hsdk/pkg/env/v1 -v -cover
go test github.com/hellohippo/hsdk/pkg/tools/v1 -v -cover
