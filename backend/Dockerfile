FROM golang:latest AS builder

WORKDIR /app 

COPY . .

RUN cd cmd/ && go build -o rss_backend


FROM ubuntu:latest

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /app/cmd/rss_backend .
COPY --from=builder /app/cmd/config.toml .

CMD ["./rss_backend", "-f", "config.toml"]