#!/bin/bash

## Uses `scss` to compile the stylesheets in `web/stylesheets`
## Requires SCSS available on the path

set -e
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

sass --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.css
sass --style=compressed --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.min.css
