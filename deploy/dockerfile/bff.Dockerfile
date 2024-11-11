FROM golang:1.23 AS builder

COPY . /go/src
WORKDIR /go/src

ENV GO111MODULE on
ENV GOPATH /go
ENV GOPROXY https://goproxy.io,direct

ENV GOSUMDB off
ENV CGO_ENABLED 0

RUN go mod tidy
RUN go build -o bff -ldflags="-s -w"  ./app/bff/cmd/bff/

FROM debian:buster
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
RUN apt-get -y install --no-install-recommends ca-certificates curl
COPY --from=builder /go/src/bff /
WORKDIR /
ENV TZ Asia/Shanghai
EXPOSE 88
EXPOSE 99
ENV PARAMS=""
##ENTRYPOINT
CMD ["sh","-c","mv /data/bff.log /data/bff.$(date '+%Y%m%d%H%M%S').log || true && /bff $PARAMS >/data/bff.log 2>&1"]