FROM golang:latest AS builder

WORKDIR /app 

COPY . .

RUN cd cmd/ && go build -o rss_backend


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/cmd/rss_backend .
COPY --from=builder /app/cmd/config.toml .

CMD ["./rss", "-f", "config.toml"]