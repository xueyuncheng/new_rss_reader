FROM golang:latest AS builder

WORKDIR /app 

COPY . .

RUN cd cmd/ && go build -o rss_backend


FROM ubuntu:latest

# 在安装tzdata之前最好先update和upgrade，以防apt-get获取不到tzdata
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN apt install -y ca-certificates

WORKDIR /app

COPY --from=builder /app/cmd/rss_backend .
COPY --from=builder /app/cmd/config.toml .

CMD ["./rss_backend", "-f", "config.toml"]