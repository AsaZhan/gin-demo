FROM golang:1.18.2-alpine3.16 as builder

COPY . /gin-demo
WORKDIR /gin-demo
RUN rm -rf /gin-demo/log
RUN go build --mod=vendor -ldflags "-s -w" -o gin-demo server/main.go

FROM alpine:3.16
COPY --from=builder /gin-demo/gin-demo /usr/local/bin/
RUN chmod 755 /usr/local/bin/gin-demo

CMD ["gin-demo"]