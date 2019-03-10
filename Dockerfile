FROM golang:1.12.0-alpine AS build-env

WORKDIR /go/src/bashserver
COPY . .

RUN go install
RUN ls -l
RUN find / -name bashserver

FROM alpine

RUN apk add --update curl && rm -rf /var/cache/apk/*
COPY --from=build-env /go/bin/bashserver /usr/bin/bashserver

EXPOSE 8080

ENTRYPOINT [ "bashserver" ]
