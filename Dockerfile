FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.14.0
COPY --from=builder /app .
EXPOSE 80
CMD ["./main"]