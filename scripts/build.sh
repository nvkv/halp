#!/usr/bin/env bash

export GO111MODULE=on

BUILD_DIR="$PWD/build"
GOPACKAGE_NAME="github.com/nvkv/halp"

mkdir -p "$BUILD_DIR"
cd "$BUILD_DIR"

VERSION=${HALP_VERSION:-alpha}

TARGET_GOOSES=("linux" "darwin")

# No need to cross-compile it for docker anyway
if [ ! -z ${1} ] && [ ${1} == "linux" ]; then
		TARGET_GOOSES=("linux")
fi

for OS in ${TARGET_GOOSES[@]}
do
		echo "Building for ${OS}"
		export GOOS=${OS}
		export GOARCH=amd64
		export CGO_ENABLED=0
		go build -o "halp-${GOOS}-${GOARCH}" -v -ldflags "-X main.VERSION=${VERSION}" "$GOPACKAGE_NAME"
done
