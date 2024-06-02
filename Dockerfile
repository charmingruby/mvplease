# Estagio de build
FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./mvplease ./cmd/api/main.go

# Estagio de execuçao
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
COPY --from=builder /build/mvplease ./mvplease
COPY .env /app
EXPOSE 3000
CMD [ "/app/mvplease" ]