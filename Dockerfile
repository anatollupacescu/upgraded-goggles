FROM golang:1.15
RUN mkdir -p /app
WORKDIR /app
ADD . /app

RUN go test -count=1 -race ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
-ldflags "-s -w" \
-o exchange ./cmd/

# leaving here the default one, for convenience
ENV EXCHANGE_RATES_URL https://api.exchangeratesapi.io/latest?base=EUR

CMD ["./exchange"]
