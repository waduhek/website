FROM --platform=$BUILDPLATFORM golang:1.25.5 AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./build/website ./cmd/website

FROM --platform=$BUILDPLATFORM scratch AS runner
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/build/website .
ENTRYPOINT ["/website"]
