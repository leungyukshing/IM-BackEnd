#!/usr/bin/env bash

rm -rf cover
mkdir cover

feat="feat:[[:space:]].*$"
fix="fix:[[:space:]].*$"
test="test:[[:space:]].*$"
chore="chore:[[:space:]].*$"
doc="doc:[[:space:]].*$"

commit_msg_reg=(
  $feat
  $fix
  $test
  $chore
  $doc
)

# check commit message
check_commit_msg () {
  for regex in ${commit_msg_reg[*]}
    do
      if [[ $TRAVIS_COMMIT_MESSAGE =~ $regex ]]; then
          printf "do match \n\n" "$TRAVIS_COMMIT_MESSAGE" "$regex ✅"
          return 0
      else
        printf "does not match \n\n" "$TRAVIS_COMMIT_MESSAGE" "$regex ❗"
      fi
    done
  return 1
}

#check_commit_msg
#if [[ $? == 1 ]]
#then
#  printf "\n\n" "$TRAVIS_COMMIT_MESSAGE" "commit message failed match "
#  exit 1
#fi

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

# check go vet
printf "3. go vet checking"
go vet ./...
printf "go vet passed!"

# go install loop check
go get github.com/evanj/loopcheck
go install github.com/evanj/loopcheck

# run loop check
printf "4. loop check"
godir=$(go list ./... | grep -v vendor)
for directory in $godir
do
  gofiles=$(ls -1 ${GOPATH}/src/${directory}/*.go)
  for gofile in $gofiles
  do
    printf "start running loop check on ${gofile}"
    ${GOPATH}/bin/loopcheck ${gofile}
    if [[ $? == 1 ]]
    then
      printf "loopcheck failed on ${gofile}"
      exit 1
    fi
    printf "done loopcheck on ${gofile}"
  done
done
printf "loop check passed! ✅"

# run test
godir=$(go list ./... | grep -v vendor)
for directory in $godir
do
  printf "start running test on $directory"
  go test -v -parallel 1 $directory -coverprofile=coverage.out
  printf "done test on $directory"
done

go tool cover -func=coverage.out
go tool cover -html=coverage.out -o cover/database.html