
#!/bin/bash

## Builds all the templates using hero, skipping if unchanged

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

function tmpl {
  fsrc="tmp/$1.hashcode"
  ftgt="tmp/$1.hashcode.tmp"

  if [ ! -d "gen/$1" ]; then
    rm -f "$fsrc"
  fi

  find -s "$2" -type f -exec md5sum {} \; | md5sum > "$ftgt"

  if cmp -s "$fsrc" "$ftgt"; then
    rm "$ftgt"
  else
    echo "updating [$2] templates"
    mv "$ftgt" "$fsrc"
    rm -rf gen/$1
    hero -extensions .html,.sql -source "$2" -pkgname $1 -dest gen/$1
  fi
}

tmpl "query" "query/sql"
tmpl "transcripttemplates" "web/transcript"
tmpl "templates" "web/templates"
tmpl "admintemplates" "web/admin"
