FROM golang:1.26.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -buildvcs=false -o /app/bin/api ./cmd/api

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache ca-certificates
RUN mkdir -p /app/uploads

COPY --from=builder /app/bin/api /app/api
COPY migrations /app/migrations

CMD ["/app/api"]
