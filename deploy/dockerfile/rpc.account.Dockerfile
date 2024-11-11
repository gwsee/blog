FROM golang:1.23 AS builder

COPY . /go/src
WORKDIR /go/src

ENV GO111MODULE on
ENV GOPATH /go
ENV GOPROXY https://goproxy.io,direct

ENV GOSUMDB off
ENV CGO_ENABLED 0

RUN go mod tidy
RUN go build -o account -ldflags="-s -w"  app/account/cmd/account/main.go

FROM golang:1.23 AS builder

# 设置工作目录为项目根目录
WORKDIR /app

# 复制 Go 模块文件，并下载依赖
COPY go.mod go.sum ./
RUN go mod tidy

# 复制剩余的项目文件
COPY . .

# 构建 Go 应用
RUN go build -o account -ldflags="-s -w" ./app/account/cmd/main.go

FROM debian:buster
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
RUN apt-get -y install --no-install-recommends ca-certificates curl
COPY --from=builder /go/src/account /
WORKDIR /
ENV TZ Asia/Shanghai
EXPOSE 88
EXPOSE 99
ENV PARAMS=""
##ENTRYPOINT
CMD ["sh","-c","mv /data/account.log /data/account.$(date '+%Y%m%d%H%M%S').log || true && /account $PARAMS >/data/account.log 2>&1"]