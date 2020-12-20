# compile nnm
FROM golang:alpine as builder
ENV LANG C.UTF-8

COPY / /src/nnm
RUN set -ex \
 && apk --no-cache add \
      build-base \
 && cd /src/nnm \
 && go build -v

#build container
FROM alpine

ENV LANG C.UTF-8

LABEL maintainer "AveryanAlex <averyanalex@gmail.com>"


COPY --from=builder /src/nnm/nnm  /app/bin
RUN set -ex \
 && mkdir -p /app/{config,public,storage}

ENTRYPOINT ["/app/bin/nnm"]
CMD ["start"]
WORKDIR /app
