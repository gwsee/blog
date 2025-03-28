FROM golang:1.23 AS builder
#COPY ../.. /go/src
# 设置环境变量
ENV GO111MODULE on
ENV GOPATH /go
ENV GOPROXY https://proxy.golang.org,direct
ENV GOSUMDB off
ENV CGO_ENABLED 0

# 优先复制依赖声明文件（利用Docker层缓存）
WORKDIR /go/src
COPY go.mod go.sum ./
# 如果存在vendor目录，复制vendor（确保本地已执行go mod vendor）
COPY vendor ./vendor
## 下载依赖（仅在go.mod/go.sum变更时执行） ##本地能够成功运行就可以不需要
#RUN go mod download
# 复制项目代码（放在最后，避免代码变更导致缓存失效）
COPY . .
RUN go build -mod=vendor -o palaces -ldflags="-s -w" ./app/palaces/cmd/palaces/

FROM registry.cn-hangzhou.aliyuncs.com/pf94514203/yunchuang:base
COPY --from=builder /go/src/palaces /
WORKDIR /
ENV TZ Asia/Shanghai
EXPOSE 88
EXPOSE 99
ENV PARAMS=""
##ENTRYPOINT
CMD ["sh","-c","mv /data/palaces.log /data/palaces.$(date '+%Y%m%d%H%M%S').log || true && /palaces $PARAMS >/data/palaces.log 2>&1"]