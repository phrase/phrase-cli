#!/bin/bash
set -e

SRC_DIR=$GOPATH/src/github.com/phrase/phrase-cli
mkdir -p $SRC_DIR
cd $SRC_DIR
tar --extract
BIN_DIR=$(mktemp -d)
bash ./build/go_build.sh $BIN_DIR
cd $BIN_DIR
tar --create .
