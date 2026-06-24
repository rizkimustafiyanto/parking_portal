# Build Stage
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags="-s -w" -o api ./cmd/api

# Runtime Stage
FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/api /app/api
COPY --from=builder /app/migrations /app/migrations

RUN mkdir -p /app/uploads

ENV TZ=Asia/Jakarta

EXPOSE 8100

CMD ["/app/api"]