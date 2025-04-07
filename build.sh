#!/bin/bash
set -e

echo "Building jokeapp CLI..."

go build -o jokeapp main.go config.go prompt.go

echo "Build complete. Binary created: ./jokeapp"