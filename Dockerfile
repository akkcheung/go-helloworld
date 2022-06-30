FROM golang:1.18-alpine3.16 AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/go-helloworld /app/ 
COPY views/ /app/views
COPY assets/ /app/assets
WORKDIR /app
CMD ["./go-helloworld"]
