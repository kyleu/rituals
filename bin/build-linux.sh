SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

arch=amd64
os=linux

echo "Building [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release
mkdir -p ./build/$os/$arch
mv ./build/release/rituals ./build/$os/$arch/rituals
