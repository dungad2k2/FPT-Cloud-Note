FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/monitor/

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/server /server

EXPOSE 8880

CMD ["/server"]

