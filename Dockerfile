FROM golang:1.25.2 AS builder
WORKDIR /app
COPY . .
RUN make build

FROM golang:1.25.2 AS runner
WORKDIR /app
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/build/website .
CMD ["/app/website"]
