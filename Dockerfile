FROM golang:1.24-alpine AS builder

WORKDIR /tmp/build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go

FROM alpine:latest
COPY --from=builder /tmp/build/app /usr/bin/app

CMD ["app"]