FROM golang:1.14.4-alpine3.12 AS builder


ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/mon_tpx
COPY . .

RUN go mod init && go get && go build

FROM alpine:3.12.0

COPY --from=builder /go/src/mon_tpx/mon_tpx /bin/

CMD ["mon_tpx"]
