FROM golang:1.18.3-alpine3.16 AS build_base

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY .. .

RUN go build -o  carrier-controller  ./

# Start fresh from a smaller image
FROM alpine:3.16

WORKDIR /app

COPY --from=build_base /app/carrier-controller  /app/carrier-controller
COPY hack/config /app/

RUN chmod 777  /app/carrier-controller

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./carrier-controller"]