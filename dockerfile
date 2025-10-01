# Stage 1: build
FROM golang:1.23 AS builder
WORKDIR /app

# cache modules
COPY go.mod go.sum ./
RUN go mod download

# copy sources & build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/bin/app .

# Stage 2: runtime
FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/bin/app .
COPY config.env .
EXPOSE 8080
ENV PORT=8080
CMD ["./app"]
