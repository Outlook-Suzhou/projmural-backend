FROM golang:1.18.3-bullseye
WORKDIR $GOPATH/src/projmural
ADD . ./
RUN go build -o projmural .
ENTRYPOINT  ["./projmural"]