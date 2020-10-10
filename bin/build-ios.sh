#!/bin/bash

## Builds the project as an iOS framework and builds the native app in `projects/ios`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Building [ios]..."

mkdir -p build/ios/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/ios/RitualsServer.framework -target=ios github.com/kyleu/rituals.dev/lib
bin/asset-reset.sh

echo "Building [ios] app..."

cd projects/ios/rituals

xcodebuild -project rituals.xcodeproj

cd build/Release-iphoneos/

cp -R rituals.app ../../../../../build/ios
