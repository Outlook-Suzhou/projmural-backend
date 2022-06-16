FROM golang:1.18.3-bullseye
WORKDIR $GOPATH/src/projmural
ADD . ./
ENV GOPROXY https://goproxy.cn
RUN go build -o projmural .
ENTRYPOINT  ["./projmural","--env=docker"]