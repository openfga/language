FROM docker.io/library/eclipse-temurin:25@sha256:ce2ef9d2d67a68de5e196d4de9d743327c44a10b57fe02ec39f572c652cacfd3

ARG ANTLR_VERSION=4.13.1
ENV CLASSPATH .:/antlr-$ANTLR_VERSION-complete.jar:$CLASSPATH
ADD https://www.antlr.org/download/antlr-$ANTLR_VERSION-complete.jar /antlr.jar
RUN chmod +r /antlr.jar

WORKDIR /app

ENTRYPOINT ["java", "-jar", "/antlr.jar"]
