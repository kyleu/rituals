SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

./bin/build-client.sh
./bin/build-css.sh

./bin/build-macos.sh
./bin/build-linux.sh
./bin/build-linux-arm.sh
./bin/build-windows.sh

mkdir -p ./build/release/stage

cd build/release/stage

# cp -R "$dir/data" ./data

cp "$dir/build/linux/amd64/npn" ./npn
zip -r "$dir/build/release/npn.linux.zip" *
rm ./npn

cp "$dir/build/linux/arm64/npn" ./npn
zip -r "$dir/build/release/npn.linux.arm.zip" *
rm ./npn

cp "$dir/build/darwin/amd64/npn" ./npn
zip -r "$dir/build/release/npn.macos.zip" *
rm ./npn

cp "$dir/build/windows/amd64/npn.exe" ./npn.exe
zip -r "$dir/build/release/npn.windows.zip" *
rm ./npn.exe

rm -rf "$dir/build/release/stage"
