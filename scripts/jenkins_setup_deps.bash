#!/bin/bash
set +x
set -e

source ./scripts/jenkins_env.bash

mkdir -p ./goworkspace/bin
mkdir -p ./goworkspace/src/github.com/fuserobotics
ln -fs $(pwd) ./goworkspace/src/github.com/fuserobotics/gogame

pushd ${GOPATH}/src/github.com/fuserobotics/gogame
go get ./...
popd