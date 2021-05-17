FROM golang:1.16.4-alpine as builder

ADD . /usedwords/
WORKDIR /usedwords/

RUN go build -o usedwords main.go

FROM alpine
RUN mkdir /app

COPY --from=builder /usedwords/usedwords /app

WORKDIR /app

RUN chmod +x /app/usedwords
RUN adduser -D -g '' uw

USER uw

ENTRYPOINT ["/app/usedwords"]