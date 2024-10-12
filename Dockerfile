FROM golang:alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /build
ADD . .
RUN go build -o main


FROM alpine
WORKDIR /app
COPY --from=builder /build/main /app
RUN apk add tzdata