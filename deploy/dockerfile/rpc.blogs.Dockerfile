FROM golang:1.23 AS builder

COPY . /go/src
WORKDIR /go/src

ENV GO111MODULE on
ENV GOPATH /go
ENV GOPROXY https://goproxy.io,direct

ENV GOSUMDB off
ENV CGO_ENABLED 0

RUN go mod tidy
RUN go build -o blog -ldflags="-s -w"  app/blogs/cmd/blogs/main.go

FROM debian:buster
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
RUN apt-get -y install --no-install-recommends ca-certificates curl
COPY --from=builder /go/src/blog /
WORKDIR /
ENV TZ Asia/Shanghai
EXPOSE 88
EXPOSE 99
ENV PARAMS=""
##ENTRYPOINT
CMD ["sh","-c","mv /data/blog.log /data/blog.$(date '+%Y%m%d%H%M%S').log || true && /blog $PARAMS >/data/blog.log 2>&1"]