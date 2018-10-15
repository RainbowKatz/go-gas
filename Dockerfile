FROM golang:alpine

ENV GOPATH=/go
ENV APP_DIR=$GOPATH/src/github.com/RainbowKatz/go-gas
ENV GOOS=darwin
ENV GOARCH=amd64

WORKDIR $APP_DIR

COPY . $APP_DIR

ENTRYPOINT go build -o build/app .

# Run with following:
# docker run --name gogas --rm -v build:/go/src/github.com/RainbowKatz/go-gas/build gogas:latest
