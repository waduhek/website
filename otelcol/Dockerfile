#
# SSL certificates
#
FROM --platform=$BUILDPLATFORM alpine:3.23 AS certs
RUN apk --update add ca-certificates

#
# Builder
#
FROM --platform=$BUILDPLATFORM golang:1.26.0 AS build-stage
ARG TARGETOS
ARG TARGETARCH
WORKDIR /build

COPY ./builder-config.yaml builder-config.yaml

RUN --mount=type=cache,target=/root/.cache/go-build \
    GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@v0.142.0
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 builder --ldflags="-w -s" --config builder-config.yaml

#
# Runner
#
FROM --platform=$BUILDPLATFORM scratch

ARG USER_UID=10001
USER ${USER_UID}

COPY ./collector-config.yaml /collector-config.yaml
COPY --from=certs /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 --from=build-stage /build/build /

ENTRYPOINT ["/website-otelcol"]
CMD ["--config", "/collector-config.yaml"]

EXPOSE 4317 4318
