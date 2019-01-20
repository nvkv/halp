#!/usr/bin/env bash

export GO111MODULE=on

#go get
go test github.com/nvkv/halp/pkg/types/data/v1 -v -cover
