#!/bin/bash

## Packages the build output for Github Releases

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..
pdir="$( pwd )"

# ./bin/build-all.sh

mkdir -p ./build/stage

rm -rf ./build/package
mkdir -p ./build/package

cd "$pdir/build/stage"

# cp -R "$pdir/data" ./data

pkg () {
  echo "$4 ($2)..."
  cp "$pdir/build/$1/$2/$3" "./$3"

  if [ $2 = "amd64" ]; then
    zip -r "$pdir/build/package/rituals.server.$4.zip" *
  else
    zip -r "$pdir/build/package/rituals.server.$4.$2.zip" *
  fi

  rm "./$3"
}

# macOS
pkg darwin amd64 rituals macos

echo "macOS app..."
cd ../darwin
zip -r "rituals.app.macos.zip" rituals.app
mv "rituals.app.macos.zip" "../package"
cd ../stage

# Linux
pkg linux amd64 rituals linux
pkg linux 386 rituals linux
pkg linux arm64 rituals linux
pkg linux arm rituals linux
pkg linux mips rituals linux
pkg linux riscv64 rituals linux

# FreeBSD
pkg freebsd amd64 rituals freebsd
pkg freebsd 386 rituals freebsd
pkg freebsd arm64 rituals freebsd
pkg freebsd arm rituals freebsd

# Windows
pkg windows amd64 rituals.exe windows
pkg windows 386 rituals.exe windows
pkg windows arm rituals.exe windows

# Docker
echo "docker..."
cp "$pdir/build/docker/rituals.docker.tar.gz" "$pdir/build/package/rituals.server.docker.tar.gz"

# WASM
echo "wasm..."
cp "$pdir/build/js/wasm/rituals.wasm" ./rituals.wasm
zip -r "$pdir/build/package/rituals.server.wasm.zip" *
rm ./rituals.wasm

# HTML
echo "html..."
pwd
cd "$pdir/projects/wasm/assets"
zip -r "$pdir/build/package/rituals.server.html.zip" *
cd "$pdir/build/stage"

# Android
echo "android library..."
cp "$pdir/build/android/rituals.aar" ./rituals.aar
zip -r "$pdir/build/package/rituals.library.android.zip" rituals.aar
rm ./rituals.aar

echo "android app..."
cp "$pdir/build/android/rituals.apk" ./rituals.apk
zip -r "$pdir/build/package/rituals.app.android.zip" rituals.apk
rm ./rituals.apk

# iOS
echo "ios framework..."
cp  -r "$pdir/build/ios/NpnServer.framework" ./NpnServer.framework

cd NpnServer.framework
rm -rf Headers
rm -rf Modules
rm -rf NpnServer
rm -rf Resources
cp -R Versions/A/* .
rm -rf Versions
cd ..

zip -r "$pdir/build/package/rituals.library.ios.zip" *
rm  -rf ./NpnServer.framework

echo "ios app..."
cd ../ios
zip -r "rituals.app.ios.zip" rituals.app
mv "rituals.app.ios.zip" "../package"
cd ../stage

rm -rf "$pdir/build/stage"
