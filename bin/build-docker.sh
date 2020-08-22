#!/bin/bash

## Makes a release build, builds a docker image, then exports and zips the output
## Requires docker

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

docker build -m 4g -t fevo-tech/rituals.dev .

mkdir -p build/docker
docker save -o build/docker/docker.tar fevo-tech/rituals.dev
cd build/docker/
rm -f docker.tar.gz
gzip docker.tar