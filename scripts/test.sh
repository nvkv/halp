#!/usr/bin/env bash

export GO111MODULE=on

go test github.com/nvkv/halp/pkg/types/data/v1 -v -cover
go test github.com/nvkv/halp/pkg/config/v1 -v -cover
