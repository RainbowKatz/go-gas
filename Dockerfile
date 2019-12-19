FROM golang:alpine

ENV APP_DIR=/app/go-gas
ENV GOOS=darwin
ENV GOARCH=amd64

WORKDIR $APP_DIR

COPY go.mod go.mod

RUN go mod download

COPY . .

# main ENTRYPOINT that outputs binary to directory that is exposed as volume in local folder outside container
ENTRYPOINT go build -o build/app .

# Build/Run with following:
# docker build -t gogas:latest . && docker run --name gogas --rm -v `pwd`/build:/app/go-gas/build gogas:latest
