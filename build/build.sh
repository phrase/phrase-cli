#!/bin/bash
set -e

wd=$(realpath $(dirname $0)/..)
source ${wd}/build/config.sh

export DIST_DIR=dist
rm -rf $DIST_DIR
mkdir $DIST_DIR

tar --create . | docker run --rm --env VERSION=${VERSION} -i golang:$GOVERSION bash -c "$(cat build/docker_build.sh)" > ${DIST_DIR}/build.tar

cd $DIST_DIR

tar --extract --file=build.tar
rm -f build.tar

# Homebrew - binary must be called phrase, because the binary name inside
# the tar will be made available system wide
for name in phrase_macosx_amd64 phrase_macosx_arm64; do
  cp $name phrase
  tar --create --mtime="@${SOURCE_DATE_EPOCH}" phrase | gzip -n > ${name}.tar.gz
  rm phrase
done

for name in phrase_linux_386 phrase_linux_amd64 phrase_linux_arm64; do
  tar --create --mtime="@${SOURCE_DATE_EPOCH}" $name | gzip -n > ${name}.tar.gz
done

if ! which zip > /dev/null; then
  echo "zip not installed"
fi

zip phrase_windows_amd64.exe.zip phrase_windows_amd64.exe > /dev/null

echo "Last change: ${LAST_CHANGE}"
echo "Version:     ${VERSION}"
