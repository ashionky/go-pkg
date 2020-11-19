FROM golang:latest
MAINTAINER  pibing

WORKDIR /app
COPY ./ /app

ENV GO111MODULE=on \
    GOPROXY="https://goproxy.io"
RUN go build

EXPOSE 9999
CMD ["./go-pkg"]

