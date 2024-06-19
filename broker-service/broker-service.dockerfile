# base go image
FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app
COPY all_key.txt /all_key.txt
CMD [ "/app/brokerApp" ]