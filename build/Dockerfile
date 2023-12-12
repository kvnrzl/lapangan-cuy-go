FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY ../ /app
RUN go mod tidy
RUN go build -o binary
ENTRYPOINT ./binary


