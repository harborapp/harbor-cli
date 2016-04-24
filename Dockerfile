FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/harbor-cli /usr/bin/
ENTRYPOINT ["/usr/bin/harbor-cli"]
CMD ["help"]
