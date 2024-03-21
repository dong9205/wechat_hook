## Build
FROM golang:1.22.0-alpine3.19 AS build
COPY . /app
WORKDIR /app
ENV GOPROXY="https://goproxy.cn,direct"
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk add git && go mod tidy && go build -o app .

## Deploy
FROM alpine:3.19
WORKDIR /
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk add tzdata curl
COPY --from=build /app/app /app
EXPOSE 9200
ENTRYPOINT ["/app"]
