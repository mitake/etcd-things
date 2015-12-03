#! /bin/bash

export GOPATH=`pwd`

go get github.com/mitake/etcd/client
pushd src/github.com/mitake/etcd/client
git checkout -b prioritize-leader origin/prioritize-leader
popd
go get golang.org/x/net/context

go build
