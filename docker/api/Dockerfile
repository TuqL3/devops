FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/config/config.yaml ./config/

USER nonroot:nonroot

EXPOSE 8080

CMD ["./app"]