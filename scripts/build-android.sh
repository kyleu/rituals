#!/bin/bash

## Uses `tsc` to compile the scripts in `client`
## Requires tsc available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
project_dir=${dir}/..
cd $project_dir

mkdir -p build/android/
gomobile bind -o build/android/rituals.aar -target=android github.com/kyleu/rituals.dev/lib
