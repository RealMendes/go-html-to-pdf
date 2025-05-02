FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY templates ./templates
COPY pdfs ./pdfs
EXPOSE 8080
CMD ["./app"] 