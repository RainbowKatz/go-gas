#!/bin/sh
#
# Run this script from repo root

mkdir build || true

# Clean old binary
rm ./build/app || true

# 1) Rebuild Docker image to update code
docker build -t gogas:latest .

# 2) Run container to generate new binary
docker run --rm -v $(pwd)/build:/go/src/github.com/RainbowKatz/go-gas/build gogas:latest
