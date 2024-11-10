FROM golang:1.23 AS builder

COPY . /
WORKDIR /

ENV GO111MODULE on
ENV GOPATH /go
ENV GOPROXY https://goproxy.io,direct

ENV GOSUMDB off
ENV CGO_ENABLED 0

RUN go mod tidy
RUN go build -o account -ldflags="-s -w"  app/account/cmd/main.go

FROM debian:buster
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
RUN apt-get -y install --no-install-recommends ca-certificates curl
COPY --from=builder /account /
WORKDIR /
ENV TZ Asia/Shanghai
EXPOSE 88
EXPOSE 99
ENV PARAMS=""
##ENTRYPOINT
CMD ["sh","-c","mv /data/account.log /data/account.$(date '+%Y%m%d%H%M%S').log || true && /account $PARAMS >/data/account.log 2>&1"]