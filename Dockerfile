# compile nnm
FROM golang:alpine as builder
ENV LANG C.UTF-8

COPY / /src
RUN set -ex \
 && ls -a / \
 && apk --no-cache add \
      build-base \
 && cd /src \
 && go build -v

#build container
FROM alpine

ENV LANG C.UTF-8

LABEL maintainer "AveryanAlex <averyanalex@gmail.com>"


COPY --from=builder /src/main  /usr/bin/
RUN set -ex

CMD [ "/bin/main" ]
