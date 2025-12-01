FROM golang:1.25.4 AS builder
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./build/website ./cmd/website

FROM scratch AS runner
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/build/website .
ENTRYPOINT ["/website"]
