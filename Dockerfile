ARG GO_VERSION=1.22.1-alpine3.18

FROM golang:${GO_VERSION} as builder

WORKDIR /app
COPY . .

WORKDIR /app

RUN go mod tidy
RUN go build .

#
FROM golang:${GO_VERSION}
WORKDIR /app

COPY --from=builder /app/wecahtdemo .

EXPOSE 8080

CMD './wecahtdemo'