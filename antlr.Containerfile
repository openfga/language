FROM docker.io/library/eclipse-temurin:24@sha256:7493205ffe6caa8074fa8a06a276bb1c5ac41d3dd0fd43a0db66d7f776e80b3e

ARG ANTLR_VERSION=4.13.1
ENV CLASSPATH .:/antlr-$ANTLR_VERSION-complete.jar:$CLASSPATH
ADD https://www.antlr.org/download/antlr-$ANTLR_VERSION-complete.jar /antlr.jar
RUN chmod +r /antlr.jar

WORKDIR /app

ENTRYPOINT ["java", "-jar", "/antlr.jar"]
