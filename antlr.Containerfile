FROM docker.io/library/eclipse-temurin:25@sha256:2d5d5e9b533342a565907c84f47a2ec298c40cde7b13a3ca85d4950fa456a641

ARG ANTLR_VERSION=4.13.1
ENV CLASSPATH .:/antlr-$ANTLR_VERSION-complete.jar:$CLASSPATH
ADD https://www.antlr.org/download/antlr-$ANTLR_VERSION-complete.jar /antlr.jar
RUN chmod +r /antlr.jar

WORKDIR /app

ENTRYPOINT ["java", "-jar", "/antlr.jar"]
