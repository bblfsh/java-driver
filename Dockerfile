FROM alpine:3.5

ARG RUNTIME_NATIVE_VERSION
ENV RUNTIME_NATIVE_VERSION $RUNTIME_NATIVE_VERSION

RUN apk add --no-cache openjdk8-jre="$RUNTIME_NATIVE_VERSION"

WORKDIR /opt/driver/bin/

ADD native/target/native-jar-with-dependencies.jar /opt/driver/bin/
RUN echo "#!/bin/sh" > native && \
    echo "java -jar native-jar-with-dependencies.jar" > native && \
    chmod +x native

CMD /opt/driver/bin/native