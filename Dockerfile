FROM golang:1.16-alpine AS builder
WORKDIR /app/
COPY . .
RUN apk --no-cache add tzdata \
&& go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct \
&& go mod download
RUN go build -o main

FROM alpine:latest
WORKDIR /cmd/
COPY --from=builder /router/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["./main"]