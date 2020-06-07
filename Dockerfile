FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o sucursales .


FROM alpine

WORKDIR /app

COPY --from=builder /app/sucursales .

ENTRYPOINT ./sucursales