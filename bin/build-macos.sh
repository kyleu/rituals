#!/bin/bash

## Builds the project as a macOS server and builds the native app in `projects/macos`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Building [macos] app..."

mkdir -p build/darwin/

bin/build.sh darwin amd64
cp build/darwin/amd64/rituals projects/macos/rituals/rituals/rituals-server

cd projects/macos/rituals

xcodebuild -project rituals.xcodeproj

cd build/Release/

cp -R rituals.app ../../../../../build/darwin
