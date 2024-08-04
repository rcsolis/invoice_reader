FROM golang:1.22-buster as builder

WORKDIR /app
# Copy source files
COPY . .

# Build
RUN make build

FROM debian:buster-slim as release

WORKDIR /app

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/ .
RUN chmod +x main

CMD ["/app/main"]