#!/bin/sh

cd /Users/katzt007/git/me/go-gas

# Clean old binary
rm ./build/app || true

# 1) Rebuild Docker image to update code
docker build -t gogas:latest .

# 2) Run container to update binary
docker run --rm -v /Users/katzt007/git/me/go-gas/build:/go/src/gogas/gogas/build gogas:latest

# 3) Execute binary
echo ""
echo ""
echo "Running app..."
echo ""
./build/app