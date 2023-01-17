FROM alpine:latest AS phrase-cli

ARG TARGETOS
ARG TARGETARCH

COPY dist/${TARGETOS}/${TARGETARCH} /usr/bin/phrase

ENTRYPOINT ["/usr/bin/phrase"]
