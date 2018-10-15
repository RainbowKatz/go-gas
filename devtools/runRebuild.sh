#!/bin/sh
#
# Run this script from repo root

# 1) Run rebuild script to generate new binary
sh ./devtools/rebuild.sh

# 2) Execute binary
echo ""
echo ""
echo "Running app..."
echo ""
./build/app