#!/usr/bin/env bash

export GO111MODULE=on

#go get
go test github.com/nvkv/halp/pkg/types/v1 -v -cover
