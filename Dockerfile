# compile nnm
FROM golang:alpine as builder
ENV LANG C.UTF-8

RUN set -ex \
 && apk --no-cache add \
      build-base \
 && git clone "https://github.com/yggdrasil-network/yggdrasil-go.git" /src \
 && cd /src \
 && git reset --hard v${YGGDRASIL_VERSION} \
 && GOARCH=$BUILD_ARCH ./build

#build container
FROM $BUILD_FROM

ENV LANG C.UTF-8

LABEL maintainer "AveryanAlex <averyanalex@gmail.com>"


COPY --from=builder /src/yggdrasil    /usr/bin/
COPY --from=builder /src/yggdrasilctl /usr/bin/
COPY                start.sh          /
RUN chmod a+x /start.sh
RUN set -ex

CMD [ "/start.sh" ]
