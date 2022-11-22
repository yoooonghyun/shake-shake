#!/bin/bash

pushd src
swag init
go build
mv src ../out
popd
./out