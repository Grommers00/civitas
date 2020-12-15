#!/bin/bash

cd functions

for dir in *; do
    echo "Building $dir Handler"
    cd ./$dir/
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
    cp $dir "../../.aws-sam/build/${dir^}Handler/"
    cd ..
done
