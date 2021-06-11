#!/usr/bin/env bash

set -e
echo "" > coverage.txt

declare -a pkgs_to_test=("pkg/nijiparser/main_test.go" 
                         "pkg/bitly/main_test.go" 
                         "pkg/ytpicker/main_test.go" 
                         "pkg/utils/main_test.go")

declare -a pkgs=("pkg/nijiparser/main.go" 
                 "pkg/bitly/main.go" 
                 "pkg/ytpicker/main.go" 
                 "pkg/utils/main.go")

for((i=0; i<${#pkgs_to_test[@]}; i++)) do
    go test -race -coverprofile=profile.out -covermode=atomic "${pkgs_to_test[i]}" "${pkgs[i]}"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done