FROM alpine:edge
MAINTAINER Thomas Boerger <thomas@webhippie.de>

RUN apk update && \
  apk add \
    ca-certificates \
    bash && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    umschlag && \
  adduser -D \
    -h /home/umschlag \
    -s /bin/bash \
    -G umschlag \
    -u 1000 \
    umschlag

COPY umschlag-cli /usr/bin/

USER umschlag
ENTRYPOINT ["/usr/bin/umschlag-cli"]
CMD ["help"]
