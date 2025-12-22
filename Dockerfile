# build stage (Deploy)
FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o intersection-app

# runtime stage
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/intersection-app .
EXPOSE 3232

CMD ["./intersection-app"]
