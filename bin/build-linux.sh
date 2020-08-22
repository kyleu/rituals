SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

arch=amd64
os=linux

echo "Building [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release
mkdir -p ./build/$os/$arch
mv ./build/release/rituals ./build/$os/$arch/rituals