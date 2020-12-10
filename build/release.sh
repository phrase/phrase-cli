#!/bin/bash

set -eo pipefail

echo "Build release $VERSION"

sed -e "s/VERSION/${VERSION}/g" ./build/innosetup/phrase-cli-386.iss.template > ./build/innosetup/phrase-cli-386.iss
sed -e "s/VERSION/${VERSION}/g" ./build/innosetup/phrase-cli.iss.template > ./build/innosetup/phrase-cli.iss

# Build client
./build/build.sh
./build/innosetup/create_installer.sh

# build docker image

IMAGE_PREFIX=phrase/phrase-cli
IMAGE=${IMAGE_PREFIX}:${VERSION}

echo build docker image ${IMAGE}
docker build --tag ${IMAGE} -f ./Dockerfile .

echo push image ${IMAGE}
docker push ${IMAGE}

# Create release
function create_release_data()
{
  cat <<EOF
{
  "tag_name": "${VERSION}",
  "name": "${VERSION}",
  "draft": true,
  "prerelease": false
}
EOF
}

echo "Create release $VERSION"
api_url="https://api.github.com/repos/phrase/phrase-cli/releases?access_token=${GITHUB_TOKEN}"
response="$(curl --data "$(create_release_data)" ${api_url})"
release_id=$(echo $response | python -c "import sys, json; print json.load(sys.stdin)['id']")

if [ -z "$release_id" ]
then
  echo "Failed to create GitHub release"
  echo $response
  exit 1
else
  echo "New release created created with id: ${release_id}"
fi

# Upload artifacts
DIST_DIR="./dist"
for file in "$DIST_DIR"/*; do
  echo "Uploading ${file}"
  asset="https://uploads.github.com/repos/phrase/phrase-cli/releases/${release_id}/assets?name=$(basename "$file")&access_token=${GITHUB_TOKEN}"
  curl -sS --data-binary @"$file" -H "Content-Type: application/octet-stream" $asset > /dev/null
  echo Hash: $(sha256sum $file)
done

echo "Publishing release"
curl \
  --silent \
  -X PATCH \
  -H "Authorization: token ${GITHUB_TOKEN}" \
  -H "Accept: application/vnd.github.v3+json" \
  "https://api.github.com/repos/phrase/phrase-cli/releases/${release_id}" \
  -d '{"draft": false}' > /dev/null

echo "Release successful"

# update homebrew-brewed
./build/update_brew.sh
