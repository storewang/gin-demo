ARG GO_VERSION=1.24.11-alpine3.23

FROM golang:${GO_VERSION} as builder

WORKDIR /app
COPY . .

WORKDIR /app

RUN go mod tidy
RUN go build .

#
FROM golang:${GO_VERSION}
WORKDIR /app

COPY --from=builder /app/wechatdemo .

EXPOSE 8080

CMD './wechatdemo'