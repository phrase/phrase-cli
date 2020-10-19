FROM alpine:latest AS phrase-cli

COPY dist/phrase_linux_amd64 /usr/bin/phrase

ENTRYPOINT ["/usr/bin/phrase"]
