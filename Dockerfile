FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

RUN go mod download

COPY . .

RUN go build -o cart ./cmd

FROM alpine

WORKDIR /app

COPY --from=builder /build/cart cart
COPY ./internal/database/migrate migrate/

CMD ["/app/cart"]
