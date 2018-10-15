FROM golang:alpine

ENV GOPATH=/go
ENV APP_DIR=$GOPATH/src/github.com/go-gas
ENV GOOS=darwin
ENV GOARCH=amd64

WORKDIR $APP_DIR

COPY . $APP_DIR

ENTRYPOINT go build -o build/app .

# Run with following:
# docker run --rm -v /Users/katzt007/git/me/go-gas/build:/go/src/repo.domain/gogas/build gogas:latest
