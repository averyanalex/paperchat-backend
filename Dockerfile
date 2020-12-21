FROM alpine

ENV LANG C.UTF-8

LABEL maintainer "AveryanAlex <averyanalex@gmail.com>"


COPY nnm-linux-x64  /usr/bin
RUN set -ex \
 && mkdir -p /app/{config,public,storage}

CMD ["nnm-linux-x64"]
WORKDIR /etc/nnm
