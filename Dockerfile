FROM openjdk:8-jdk-alpine
MAINTAINER source{d}

ADD jar /jar

CMD ["./jar/babelfish-java-driver"]
