#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Tags the git repo using the first argument or the incremented minor version

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

git fetch --tags

TGT=${1-none}
if [[ $TGT == "none" ]]; then
  TGT=$(git describe --match "v[0-9]*" --tags --abbrev=0 | sed -e 's/v//g')
  TGT=$(echo ${TGT} | awk -F. -v OFS=. '{$NF++;print}')
fi
if [[ ${TGT:0:1} == "v" ]]; then
  TGT="${TGT:1}"
fi

echo $TGT

pat="^[0-9]"
if [[ $TGT =~ $pat ]]; then
  sed -i.bak -e "s/version = \\\"[v]*[0-9]*[0-9]\.[0-9]*[0-9]\.[0-9]*[0-9]\\\"/version = \"${TGT}\"/g" ./main.go
  rm -f "./main.go.bak"
  sed -i.bak -e "s/\\\"version\\\": \\\"[v]*[0-9]*[0-9]\.[0-9]*[0-9]\.[0-9]*[0-9]\\\"/\"version\": \"${TGT}\"/g" ./.projectforge/project.json
  rm -f "./.projectforge/project.json.bak"
fi

make build

git add .

if [[ $TGT =~ $pat ]]; then
  git commit -m "v${TGT}" || true
  git tag "v${TGT}"
else
  git commit -m "${TGT}" || true
  git tag "${TGT}"
fi

git push
git push --tags