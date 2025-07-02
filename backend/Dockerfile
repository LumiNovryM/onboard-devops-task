ARG GOVERSION=1.24.4
ARG ALPINEVERSION=3.22

FROM golang:${GOVERSION}-alpine${ALPINEVERSION} as builder

WORKDIR /opt

COPY . .

ENV GOSUMDB=off
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/user

FROM alpine:${ALPINEVERSION}

RUN apk add --no-cache tzdata openssh busybox-extras postgresql-client

COPY . .

COPY --from=builder /go/bin/user /go/bin/user

ENV FIBER_MODE=release
ENV TZ=Asia/Jakarta

COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh
ENTRYPOINT ["/wait-for-postgres.sh", "/go/bin/user"]