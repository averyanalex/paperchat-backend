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

<<<<<<< HEAD
//ENTRYPOINT ["/app/bin/nnm"]
CMD ["/app/bin/nnm"]
=======
CMD ["nnm-linux-x64"]
>>>>>>> 1cd15a9e2395008742b528a1993f9845cc28c92b
WORKDIR /etc/nnm
