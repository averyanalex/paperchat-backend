# compile nnm
FROM golang:alpine as builder
ENV LANG C.UTF-8

COPY / /src/nnm
RUN set -ex \
 && apk --no-cache add \
      build-base \
 && cd /src/nnm \
 && go build -v \
 && mv nnm nnm-linux-x64

# Build container
FROM alpine

ENV LANG C.UTF-8

LABEL maintainer "AveryanAlex <averyanalex@gmail.com>"


COPY --from=builder /src/nnm/nnm-linux-x64  /usr/bin
RUN set -ex \
 && mkdir -p /app/{config,public,storage}

CMD ["nnm-linux-x64"]
WORKDIR /etc/nnm
