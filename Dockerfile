FROM golang:1.18.3 AS builder
WORKDIR /module
COPY .  .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /module/app .
EXPOSE 8000
CMD ["./app"]