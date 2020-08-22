FROM golang:1.15-alpine
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o appbin main.go

FROM alpine:3.12
WORKDIR /app
COPY --from=0 /app/appbin /app/appbin
COPY --from=0 /app/static /app/static
CMD ["/app/appbin"]
