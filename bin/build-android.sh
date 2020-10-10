#!/bin/bash

## Builds the project as an android framework and builds the native app in `projects/android`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

pdir="$( pwd )"

echo "Building [android]..."

mkdir -p build/android/

bin/asset-embed.sh
echo "gomobile..."
gomobile bind -o build/android/rituals.aar -target=android github.com/kyleu/rituals.dev/lib
bin/asset-reset.sh

cd projects/android/rituals/app/libs
rm -f rituals.aar rituals-sources.jar
cp ${pdir}/build/android/rituals.aar .
cp ${pdir}/build/android/rituals-sources.jar .

cd "${pdir}/projects/android/rituals"

gradle assembleDebug

cp "app/build/outputs/apk/debug/app-debug.apk" "${pdir}/build/android/rituals.apk"

cd "${pdir}/build/android"
