################# First stage #######################
FROM golang:1.17-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/biowar/webserver-bitcoin-price

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

################# Second stage #######################
FROM scratch

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /go/src/github.com/biowar/webserver-bitcoin-price .

EXPOSE 12321

CMD ["./main"]