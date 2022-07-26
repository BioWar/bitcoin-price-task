FROM golang:1.17-alpine

RUN apk --no-cache add ca-certificates
# RUN apk add --no-cache bash

WORKDIR /bitcoin-price-api

COPY . .

RUN go build -o /webserver-bitcoin-price .

EXPOSE 12321

CMD ["/webserver-bitcoin-price"]