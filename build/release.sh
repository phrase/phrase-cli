#!/bin/bash

set -eo pipefail

echo "Build release $VERSION"

sed -e "s/VERSION/${VERSION}/g" ./build/innosetup/phrase-cli-386.iss.template > ./build/innosetup/phrase-cli-386.iss
sed -e "s/VERSION/${VERSION}/g" ./build/innosetup/phrase-cli.iss.template > ./build/innosetup/phrase-cli.iss

# Build client
./build/build.sh
# The Windows installer is an optional extra: the plain Windows binaries are
# already built above. Don't let a failure here (e.g. wine/Docker issues on the
# runner) block the rest of the release.
if ! ./build/innosetup/create_installer.sh; then
  echo "::warning title=Windows installer::Inno Setup build failed; releasing without phrase_windows_setup*.exe"
fi

# build docker image

IMAGE_PREFIX=phrase/phrase-cli
IMAGE=${IMAGE_PREFIX}:${VERSION}
IMAGE_LATEST=${IMAGE_PREFIX}:latest

echo build docker image "${IMAGE}" and ${IMAGE_LATEST}

mkdir -p dist/linux
cp dist/phrase_linux_amd64 dist/linux/amd64
cp dist/phrase_linux_arm64 dist/linux/arm64

docker buildx build --tag "${IMAGE}" --tag ${IMAGE_LATEST} --platform linux/amd64,linux/arm64 -f ./Dockerfile --push .

echo "Artifacts built and ready in dist/ directory. GitHub Release creation handled in GitHub Actions workflow."
