FROM golang:1.15-alpine
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o appbin main.go

FROM alpine:3.13.7
WORKDIR /app
COPY --from=0 /app/appbin /app/appbin
COPY --from=0 /app/quotes.json /app/quotes.json
CMD ["/app/appbin"]
