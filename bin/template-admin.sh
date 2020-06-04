#!/bin/bash

## Builds the html admin templates using hero, skipping if unchanged

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

fsrc="tmp/admin.hashcode"
ftgt="tmp/admin.hashcode.tmp"

if [ ! -d "gen/admintemplates" ]; then
  rm "$fsrc"
fi

find -s web/admintemplates -type f -exec md5sum {} \; | md5sum > "$ftgt"

if cmp -s "$fsrc" "$ftgt"; then
  rm "$ftgt"
else
  mv "$ftgt" "$fsrc"
	rm -rf gen/admintemplates
	hero -extensions .html -source web/admintemplates -pkgname admintemplates -dest gen/admintemplates
fi
