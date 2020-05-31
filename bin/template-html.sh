#!/bin/bash

## Builds the html templates using hero, skipping if unchanged

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

fsrc="tmp/html.hashcode"
ftgt="tmp/html.hashcode.tmp"

find -s web/templates -type f -exec md5sum {} \; | md5sum > "$ftgt"

if cmp -s "$fsrc" "$ftgt"; then
  rm "$ftgt"
else
  mv "$ftgt" "$fsrc"
	rm -rf gen/templates
	hero -extensions .html -source web/templates -pkgname templates -dest gen/templates
fi
