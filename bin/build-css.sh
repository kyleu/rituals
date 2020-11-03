#!/bin/bash

## Uses `scss` to compile the stylesheets in `web/stylesheets`
## Requires SCSS available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Building SCSS..."
sass --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.css
echo "Building optimized SCSS..."
sass --style=compressed --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.min.css
