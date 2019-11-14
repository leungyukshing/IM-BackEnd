#!/usr/bin/env bash

set -ex

rm -rf cover
mkdir cover

# check gofmt
gofiles=$(git diff --diff-filter=ACM --name-only origin/master..HEAD | awk '!/vendor/ && /.go$/')
printf "2. go fmt check"
if [ "$gofiles" = "" ]; then
    printf "no go files, skip it! ✅"
else
  unformatted=$(gofmt -l $gofiles)
  if [ "$unformatted" = "" ]; then
      printf "All Go files are formatted, great job! ✅"
  else
    printf "Go files must be formatted with gofmt. Please run: go fmt $unformatted first."
    exit 1
  fi
fi