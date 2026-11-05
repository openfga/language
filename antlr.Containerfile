FROM docker.io/library/eclipse-temurin:25@sha256:83f4a7fc27df94dcc7b8de04d14cbbba8da14406f921e966969c3b09000aa507

ARG ANTLR_VERSION=4.13.1
ENV CLASSPATH .:/antlr-$ANTLR_VERSION-complete.jar:$CLASSPATH
ADD https://www.antlr.org/download/antlr-$ANTLR_VERSION-complete.jar /antlr.jar
RUN chmod +r /antlr.jar

WORKDIR /app

ENTRYPOINT ["java", "-jar", "/antlr.jar"]
