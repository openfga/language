FROM docker.io/library/eclipse-temurin:11@sha256:96e6d8e43d40476fe01d5a09c5aa639d18be78292defd761429c0151f3f0ee2f

ARG ANTLR_VERSION=4.13.1
ENV CLASSPATH .:/antlr-$ANTLR_VERSION-complete.jar:$CLASSPATH
ADD https://www.antlr.org/download/antlr-$ANTLR_VERSION-complete.jar /antlr.jar
RUN chmod +r /antlr.jar

WORKDIR /app

ENTRYPOINT ["java", "-jar", "/antlr.jar"]
