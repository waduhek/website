#
# SSL certificates
#
FROM alpine:3.22 AS certs
RUN apk --update add ca-certificates

#
# Builder
#
FROM golang:1.25.4 AS build-stage
WORKDIR /build

COPY ./builder-config.yaml builder-config.yaml

RUN --mount=type=cache,target=/root/.cache/go-build \
    GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@v0.141.0
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 builder --ldflags="-w -s" --config builder-config.yaml

#
# Runner
#
FROM scratch

ARG USER_UID=10001
USER ${USER_UID}

COPY ./collector-config.yaml /collector-config.yaml
COPY --from=certs /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 --from=build-stage /build/build /

ENTRYPOINT ["/website-otelcol"]
CMD ["--config", "/collector-config.yaml"]

EXPOSE 4317 4318
