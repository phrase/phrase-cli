#!/bin/bash

set -e

docker build -t phrase/innosetup build/innosetup/
docker run --rm --env VERSION=${VERSION} --entrypoint /bin/bash -v $(pwd):/client phrase/innosetup /client/build/innosetup/docker_build.sh
