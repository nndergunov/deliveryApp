FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ../go.mod go.sum ./

RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order-service main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/order-service .
COPY --from=builder /app/config.yaml .

EXPOSE 8085

CMD ["./order-service"]
