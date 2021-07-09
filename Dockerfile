FROM golang:1.15 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -tags bindatafs -o main .

FROM alpine:3.9

RUN adduser -D -g '' appuser
USER appuser

WORKDIR /home/appuser
COPY --from=builder /app/main .

CMD ["./main"]

