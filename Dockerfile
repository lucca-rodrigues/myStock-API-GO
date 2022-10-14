FROM golang:1.19.1-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build main.go

FROM debian:buster-slim

# RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#     ca-certificates && \
#     rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main

CMD ["./app/main"]

EXPOSE 3333