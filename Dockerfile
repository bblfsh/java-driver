FROM alpine:3.5

ARG NATIVE_RUNTIME_VERSION
ENV NATIVE_RUNTIME_VERSION $NATIVE_RUNTIME_VERSION

RUN apk add --no-cache openjdk8-jre="$NATIVE_RUNTIME_VERSION"

WORKDIR /opt/driver/bin/

ADD native/target/native-jar-with-dependencies.jar /opt/driver/bin/
RUN echo "#!/bin/sh" > native && \
    echo "java -jar native-jar-with-dependencies.jar" > native && \
    chmod +x native

CMD /opt/driver/bin/native