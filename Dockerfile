FROM golang:1.18-alpine3.15

ENV ROOT /go
ENV CGO_ENABLED 0

WORKDIR ${ROOT}

RUN apk update && apk add --no-cache git

CMD sh
