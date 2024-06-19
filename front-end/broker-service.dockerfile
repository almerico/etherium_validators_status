# base go image
FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/web

RUN chmod +x /app/brokerApp
CMD [ "/app/brokerApp" ]