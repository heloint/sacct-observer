#!/usr/bin/env bash

rm -rf ./pre-compiled-binaries && mkdir -p ./pre-compiled-binaries

# WINDOWS BINARIES
# ================
echo "Compiling for Windows with 64-bit AMD architecture..."
GOOS=windows GOARCH=amd64 go build -o ./pre-compiled-binaries/sacct-observer-amd64.exe ../cmd/sacct-observer/main.go # 64-bit (amd)

echo "Compiling for Windows with 32-bit AMD architecture..."
GOOS=windows GOARCH=386 go build -o ./pre-compiled-binaries/sacct-observer-386.exe ../cmd/sacct-observer/main.go # 32-bit (amd)

# MACOS BINARIES
# ==============
echo "Compiling for MacOS with 64-bit AMD architecture..."
GOOS=darwin GOARCH=amd64 go build -o ./pre-compiled-binaries/sacct-observer-amd64-darwin ../cmd/sacct-observer/main.go # 64-bit (amd)

echo "Compiling for MacOS with 64-bit ARM architecture..."
GOOS=darwin GOARCH=arm64 go build -o ./pre-compiled-binaries/sacct-observer-arm64-darwin ../cmd/sacct-observer/main.go # Apple Silicon (arm64)

# LINUX BINARIES
# ==============
echo "Compiling for Linux with 64-bit AMD architecture..."
GOOS=linux GOARCH=amd64 go build -o ./pre-compiled-binaries/sacct-observer-amd64-linux ../cmd/sacct-observer/main.go # 64-bit (amd)

echo "Compiling for Linux with 32-bit AMD architecture..."
GOOS=linux GOARCH=386 go build -o ./pre-compiled-binaries/sacct-observer-386-linux ../cmd/sacct-observer/main.go # 32-bit (amd)

# SHA256 CHECKSUMS
echo "Creating SHA256 Checksums..."
sha256sum ./pre-compiled-binaries/sacct-observer-386.exe | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
sha256sum ./pre-compiled-binaries/sacct-observer-amd64.exe | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
sha256sum ./pre-compiled-binaries/sacct-observer-amd64-darwin | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
sha256sum ./pre-compiled-binaries/sacct-observer-arm64-darwin | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
sha256sum ./pre-compiled-binaries/sacct-observer-amd64-linux | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
sha256sum ./pre-compiled-binaries/sacct-observer-386-linux | sed 's/\.\/pre-compiled-binaries\///g' >> ./pre-compiled-binaries/sha256_checksums.txt
