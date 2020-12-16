FROM golang:latest
MAINTAINER  pibing

WORKDIR /app
COPY . .

ENV GO111MODULE=on \
    GOPROXY="https://goproxy.io"
RUN go build

CMD ["./go-pkg"]
EXPOSE 9999

