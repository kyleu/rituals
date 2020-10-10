#!/bin/bash

## Makes a release build, builds a docker image, then exports and zips the output
## Requires docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

docker build -m 4g -t kyleu/rituals .

mkdir -p build/docker
docker save -o build/docker/rituals.docker.tar kyleu/rituals
cd build/docker/
rm -f rituals.docker.tar.gz
gzip rituals.docker.tar
