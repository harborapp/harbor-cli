FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/umschlag-cli /usr/bin/
ENTRYPOINT ["/usr/bin/umschlag-cli"]
CMD ["help"]
