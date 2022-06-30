# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /mail-app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o mail-app /mail-app/cmd/app

EXPOSE 8080

CMD [ "./mail-app" ]