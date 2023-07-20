FROM cgr.dev/chainguard/go:1.20@sha256:f4db6d6cf7edf5190b8b4923ecd8aeeeab856bc7777fcdddbd9c542aeb7aac92 AS builder

ARG BINARY_NAME=openfga-language

WORKDIR /app

COPY ../.. .
RUN CGO_ENABLED=0 go build -o ${BINARY_NAME} ./main.go

FROM cgr.dev/chainguard/static@sha256:54b589146d4dbc80a094fcbcd6b09414f3df94cde8ea6d31c44fd02692c58203

COPY --from=builder /app/${BINARY_NAME} /${BINARY_NAME}
ENTRYPOINT ["/openfga-language"]
