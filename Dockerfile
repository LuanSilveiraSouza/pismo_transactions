FROM golang:1.20 as build

RUN mkdir /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./vendor /app/vendor
COPY ./src /app/src

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -o app /app/src/main.go

#Deploy
FROM alpine:latest

WORKDIR /app

COPY --from=build /app .

ENV DB_HOSTNAME=db

ENTRYPOINT ["/app/app"]