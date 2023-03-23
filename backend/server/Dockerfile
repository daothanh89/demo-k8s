FROM golang:1.15-alpine
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o appbin main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=0 /app/appbin /app/appbin
COPY --from=0 /app/static /app/static
CMD ["/app/appbin"]
