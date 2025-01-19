FROM golang:1.23.5-alpine3.21 AS builder-env

WORKDIR /go/src/

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app /go/src/cmd/booking-server
RUN cp app /app
EXPOSE 8089
CMD ["/app"]
