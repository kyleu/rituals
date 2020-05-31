#!/bin/bash

## Builds the sql templates using hero, skipping if unchanged

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

fsrc="tmp/sql.hashcode"
ftgt="tmp/sql.hashcode.tmp"

if [ ! -d "gen/query" ]; then
  rm "$fsrc"
fi

find -s query/sql -type f -exec md5sum {} \; | md5sum > "$ftgt"

if cmp -s "$fsrc" "$ftgt"; then
  rm "$ftgt"
else
  mv "$ftgt" "$fsrc"
	rm -rf gen/query
	hero -extensions .sql -source query/sql -pkgname query -dest gen/query
fi

