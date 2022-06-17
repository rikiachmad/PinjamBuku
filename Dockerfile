FROM golang:1.17.5-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go clean --modcache
RUN go build ./backend/app/main.go

FROM alpine:3.6

WORKDIR /root/

COPY --from=builder /app/.env .
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]
