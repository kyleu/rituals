SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

./scripts/build-client.sh
./scripts/build-css.sh

for arch in amd64 386
do
  for os in darwin linux windows
  do
    echo "Building [$os $arch]..."
    env GOOS=$os GOARCH=$arch make build-release
    mkdir -p ./build/$os/$arch
    if [ "$os" = "windows" ]; then
      mv ./build/release/rituals.exe ./build/$os/$arch/rituals.exe
    else
      mv ./build/release/rituals ./build/$os/$arch/rituals
    fi
  done
done
