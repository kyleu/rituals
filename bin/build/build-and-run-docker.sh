#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Builds the Docker image and runs it

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

echo "Building [linux amd64]..."
GOOS=linux GOARCH=amd64 make build
mv ./build/debug/rituals .
docker build -t=rituals -f=./tools/release/Dockerfile .
rm ./rituals
docker run -it rituals
