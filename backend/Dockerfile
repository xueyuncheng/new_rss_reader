FROM alpine:latest

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories

RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
ENV TZ Asia/Shanghai

WORKDIR /app

COPY ./cmd/rss rss
COPY ./cmd/config.toml config.toml

CMD ["./rss", "-f", "config.toml"]