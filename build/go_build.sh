#!/bin/bash
set -e

source $(realpath $(dirname $0))/config.sh
bin_dir=$1

if [[ -z $bin_dir ]]; then
	echo "USAGE: BIN_DIR" > /dev/stderr
	exit 1
fi

function build {
	goos=$1
	goarch=$2
	name=$3
	echo "build os=${goos} arch=${goarch}" > /dev/stderr

	CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -o $bin_dir/${name} -ldflags "-X 'github.com/phrase/phrase-cli/cmd.LAST_CHANGE=${LAST_CHANGE}' -X=github.com/phrase/phrase-cli/cmd.REVISION=$REVISION -X=github.com/phrase/phrase-cli/cmd.PHRASE_CLIENT_VERSION=${VERSION} -X=github.com/phrase/phrase-cli/cmd.LIBRARY_REVISION=$LIBRARY_REVISION -extldflags '-static'" .
}

build linux   amd64   phrase_linux_amd64
build linux   386     phrase_linux_386
build darwin  amd64   phrase_macosx_amd64
build darwin  arm64   phrase_macosx_arm64
build windows amd64   phrase_windows_amd64.exe
build windows 386     phrase_windows_386.exe
