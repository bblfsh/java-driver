FROM alpine:3.7

RUN mkdir -p /opt/driver/src && \
    adduser $BUILD_USER -u $BUILD_UID -D -h /opt/driver/src

RUN apk add --no-cache make bash git curl maven openjdk8="$RUNTIME_NATIVE_VERSION"

WORKDIR /opt/driver/src
