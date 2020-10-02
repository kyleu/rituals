#!/bin/bash

## Uses `tsc` to compile the scripts in `client`
## Requires tsc available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../client

tsc --project tsconfig.json

cd $project_dir/web/assets/vendor

closure-compiler --create_source_map rituals.min.js.map rituals.js > rituals.min.js
