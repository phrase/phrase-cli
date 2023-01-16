FROM --platform=$BUILDPLATFORM alpine:latest AS phrase-cli

ARG CLI_PATH
COPY ${CLI_PATH} /usr/bin/phrase

ENTRYPOINT ["/usr/bin/phrase"]
