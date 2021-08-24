FROM golang:1.17.0-bullseye
WORKDIR $GOPATH/src/projmural
ADD . ./
RUN go build -o projmural .
ENTRYPOINT  ["./projmural"]